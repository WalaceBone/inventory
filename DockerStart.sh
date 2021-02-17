if [[ "$(docker images -q inventory 2> /dev/null)" == "" ]]; then
    docker build -t inventory .
fi
docker run --rm -it -v $(pwd):/test-inventory -w /test-inventory inventory /bin/sh
