# Gotify Notification Setup

This guide will walk you through setting up Gotify notifications. You will learn how to obtain the necessary credentials and configure them in your environment.

## 1. Get Gotify URL and App Token

First, you need to get the Gotify server URL and an app token.

1.  **Gotify Server URL**: This is the URL of your Gotify server.
2.  **App Token**: You can create a new app in the Gotify web UI and get the token.

## 2. Configure Chatz

Next, you need to configure Chatz to use the Gotify provider. You can do this by editing the `~/.chatz.ini` file or by setting environment variables.

### Using `~/.chatz.ini`

Add the following section to your `~/.chatz.ini` file:

```ini
[gotify]
PROVIDER=gotify
GOTIFY_URL=<your-gotify-server-url>
GOTIFY_TOKEN=<your-gotify-app-token>
```

### Using Environment Variables

Set the following environment variables:

```sh
export PROVIDER=gotify
export GOTIFY_URL=<your-gotify-server-url>
export GOTIFY_TOKEN=<your-gotify-app-token>
```

## 3. Send a Test Notification

Now you can send a test notification to your Gotify server.

```sh
chatz --profile=gotify "Hello from Chatz!"
```

You should receive a notification on your Gotify server.
