#!/bin/bash

NETWORK="duality-devnet"
DOCKER_COMPOSE_FILE="docker-compose.yml"
OVERRIDE_COMPOSE_FILE="./networks/${NETWORK}/docker-compose.override.yml"
ECS_PARAMS_FILE="./networks/${NETWORK}/ecs-params.yml"
OUTPUT_DIR="./compiled-services"
EXCLUDED_PROFILES="local-devnet"

# TODO: fixme
TAG=heighliner-test

ecs_params_services=$(dasel -f $ECS_PARAMS_FILE ".task_definition.services.keys().all()" | sort)
mkdir -p $OUTPUT_DIR
profiles=$(docker compose convert --profiles)


while read -r profile; do
    echo "profile: $profile"
    if [[ " $EXCLUDED_PROFILES " =~ .*\ $profile\ .* ]]; then
        echo "skipping $profile"
        break
    fi
    mkdir -p "${OUTPUT_DIR}/${profile}"
    output_file="${OUTPUT_DIR}/${profile}/docker-compose.yml"
    TAG=$TAG docker compose --profile $profile \
           -f $DOCKER_COMPOSE_FILE \
           -f $OVERRIDE_COMPOSE_FILE \
           convert > $output_file

  # ecs-cli is very particular about the format so we have to do a bit of cleanup on the generated file

    ## remove depends_on since this is handled by ecs-config.yml
    dasel delete -f $output_file '.services.all().depends_on?'

    ## convert volumes back to 'source:target' format
    volume=$(dasel -f $output_file '.services.all().volumes?.merge().first().first().join(:,source,target)')
    dasel put -f $output_file  -v $volume '.services.all().volumes?.[0]'

    ## remove volumes.*.name
    dasel delete -f $output_file '.volumes.all().name?'

    ## remove name
    dasel delete -f $output_file '.name'

    ## remove profiles

    dasel delete -f $output_file '.services.all().profiles?'

    ## remove networks

    dasel delete -f $output_file '.networks'

    echo "version: \"3\"" >> $output_file

  # Create a service specific ecs-params file
    cp $ECS_PARAMS_FILE "${OUTPUT_DIR}/${profile}"

    service_ecs_params_file="${OUTPUT_DIR}/${profile}/ecs-params.yml"

    echo $service_ecs_params_file
    ## remove all of the services not in the service specific docker-compose file
    expected_services=$(dasel -f $output_file '.services.keys().all()' | sort)
    services_to_remove=$(comm -13 <(echo "$expected_services") <(echo "$ecs_params_services"))
    echo -e "remove list: ${services_to_remove[0]} \n ****"
    while read -d ' ' unwanted_service; do
        echo "removing: $unwanted_service"
        dasel delete -f $service_ecs_params_file -o $service_ecs_params_file '.task_definition.services.'
        echo "done"
    done <<< "$services_to_remove "

done <<< "$profiles"
