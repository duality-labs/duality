
# set default values
STAGE="${STAGE:-dev}"
EB_APP_NAME="${EB_APP_NAME:-deploy-preview}"
EB_ENVIRONMENT_NAME="${EB_ENVIRONMENT_NAME:-Deploypreview-alpine-env}"
EB_BUCKET="${EB_BUCKET:-duality-elastic-beanstalk}"
EB_BUCKET_FOLDER="${EB_BUCKET_FOLDER:-deploy-preview}"

# log in to AWS Elastic Container Registry with AWS crendentials that should be in the current environment
aws ecr get-login-password --region us-west-1 | docker login --username AWS --password-stdin 172530569905.dkr.ecr.us-west-1.amazonaws.com

# Generate a fairly unique tag name for the Docker image
TAG="$STAGE-$(date +%Y-%m-%d-%s)"

docker build -t backend-hapi .
docker tag backend-hapi:latest "172530569905.dkr.ecr.us-west-1.amazonaws.com/duality-core:$TAG"
docker push "172530569905.dkr.ecr.us-west-1.amazonaws.com/duality-core:$TAG"

# Create new Elastic Beanstalk version
DOCKERRUN_STRING="{\"AWSEBDockerrunVersion\":\"1\",\"Image\":{\"Name\":\"172530569905.dkr.ecr.us-west-1.amazonaws.com/duality-core:$TAG\",\"Update\":\"true\"},\"Ports\":[{\"ContainerPort\":\"26657\"},{\"ContainerPort\":\"1317\"}],\"Logging\":\"/var/log/app\"}"
DOCKERRUN_FILE="Dockerrun-$TAG.aws.json"
echo $DOCKERRUN_STRING > $DOCKERRUN_FILE

# Upload new Elastic Beanstalk version
aws s3 cp $DOCKERRUN_FILE s3://$EB_BUCKET/$EB_BUCKET_FOLDER/$DOCKERRUN_FILE
aws elasticbeanstalk create-application-version \
  --application-name $EB_APP_NAME \
  --version-label $TAG \
  --source-bundle S3Bucket=$EB_BUCKET,S3Key=$EB_BUCKET_FOLDER/$DOCKERRUN_FILE

# Update Elastic Beanstalk environment to new version
aws elasticbeanstalk update-environment \
  --environment-name $EB_ENVIRONMENT_NAME \
  --version-label $TAG
