
# remove old containers
docker-compose down;
# build new containers
docker-compose -f docker-compose.setup.yml build;

# remove old testnet data
rm -rf .testnets
# build new testnet data
docker-compose -f docker-compose.setup.yml run --rm init-files;

# compose and run nodes together
docker-compose up --build;
