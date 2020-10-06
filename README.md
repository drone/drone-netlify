Drone plugin for deploying static websites to Netlify. _This plugin is used to publish all static websites for the Drone project, including the official documentation._

# Building

Build the plugin binary:

```text
scripts/build.sh
```

Build the plugin image:

```text
docker build -f docker/Dockerfile -t plugins/netlify .
```

# Testing

Execute the plugin from your current working directory:

```text
docker run --rm \
  -e PLUGIN_SITE=3970e0fe-8564-4903-9a55-c5f8de49fb8b \
  -e PLUGIN_PATH=./public \
  -e PLUGIN_TOKEN=your_oauth2_access_token \
  -e PLUGIN_DEBUG=true \
  -w /drone/src \
  -v $(pwd):/drone/src \
  plugins/netlify
```