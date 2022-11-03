
# set default values
STAGE="${STAGE:-dev}"
EB_APP_NAME="${EB_APP_NAME:-deploy-preview}"
EB_ENVIRONMENT_NAME="${EB_ENVIRONMENT_NAME:-duality-core-deploy-preview-env}"
EB_BUCKET="${EB_BUCKET:-duality-elastic-beanstalk-dev}"
EB_BUCKET_FOLDER="${EB_BUCKET_FOLDER:-deploy-preview}"

# log in to AWS Elastic Container Registry with AWS crendentials that should be in the current environment
aws ecr get-login-password --region us-west-1 | docker login --username AWS --password-stdin 172530569905.dkr.ecr.us-west-1.amazonaws.com

# Generate a fairly unique tag name for the Docker image
TAG="$STAGE-$(date +%Y-%m-%d-%s)"
APPLICATION_VERSION="duality-core--$TAG"

docker build -t duality-core .
docker tag duality-core:latest "172530569905.dkr.ecr.us-west-1.amazonaws.com/duality-core:$TAG"
docker push "172530569905.dkr.ecr.us-west-1.amazonaws.com/duality-core:$TAG"

# Create temporary directory to collect application version files
TMP_DIR=$(mktemp -d /tmp/$APPLICATION_VERSION.XXXXXXXX)
trap "rm -rf $TMP_DIR" EXIT

# Create new Elastic Beanstalk version
DOCKERRUN_STRING="{\"AWSEBDockerrunVersion\":\"1\",\"Image\":{\"Name\":\"172530569905.dkr.ecr.us-west-1.amazonaws.com/duality-core:$TAG\",\"Update\":\"true\"},\"Ports\":[{\"ContainerPort\":\"26657\"},{\"ContainerPort\":\"1317\"}],\"Logging\":\"/var/log/app\"}"
echo $DOCKERRUN_STRING > $TMP_DIR/Dockerrun.aws.json
cp -r .ebextensions $TMP_DIR/.ebextensions

# Create temp zip file for application update
(cd $TMP_DIR && zip -r $APPLICATION_VERSION.zip .ebextensions Dockerrun.aws.json)

# Upload new Elastic Beanstalk version
aws s3 cp $TMP_DIR/$APPLICATION_VERSION.zip s3://$EB_BUCKET/$EB_BUCKET_FOLDER/$APPLICATION_VERSION.zip
aws elasticbeanstalk create-application-version \
  --application-name $EB_APP_NAME \
  --version-label $APPLICATION_VERSION \
  --source-bundle S3Bucket=$EB_BUCKET,S3Key=$EB_BUCKET_FOLDER/$APPLICATION_VERSION.zip & \

# Update Elastic Beanstalk environment to new version
aws elasticbeanstalk update-environment \
  --environment-name $EB_ENVIRONMENT_NAME \
  --version-label $APPLICATION_VERSION
