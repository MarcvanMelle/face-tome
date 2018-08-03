# face-tome

How to use:

`docker-compose up`
`grpc_cli call 0.0.0.0:5501 facetome.FaceTome.GetNPC "npc_uuid: 'anything'"`

Working with Local Image Registry
Start by running a registry if none is currently running.
`docker run -d -p 5000:5000 --restart=always --name registry registry:2`
Then build your image
`docker build -f ./build/Dockerfile -t face-tome:latest .`
Then push your image to the registry
`docker push localhost:5000/face-tome`

Your registry is running on `localhost:5000`
To have give GKE access, you must expose port 5000 on your machine to the public internet.
We'll use Forward.
Go to `localhost:5000` in Chrome, and using the Forawrd Chrome extension, configure a DNS to use; e.g.
`face-tome-reg-adh8.fwd.wf`

Now configure your Deployment yaml to use the image `image: face-tome-reg-adh8.fwd.wf/face-tome:latest`
