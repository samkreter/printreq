# printreq

Simple container that prints the http reuqests it recieves. Used in for testing connectivty and request content. 

## Running

### Dockerhub Image

Run the image from docker hub with the below command

    docker run -p 8081:8081 pskreter/printreq:stable
    
### Building the Image

After cloning the repo, build the image with the following command

    docker build -t <yourUsername>/printreq .

## Contributing

Contributions are always welcome. Just make a pull request. It will be reviewed and merged quickly. 
