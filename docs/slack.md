# Creating and Configuring a Slack Bot Token

A Slack Bot Token is a secure way for your application to interact with Slack, allowing it to send messages, respond to events, and perform other automated tasks. This guide provides a step-by-step process to create a Slack bot, obtain a Bot Token, and configure it for use in your project.

## Step 1: Create a Slack App

1. **Visit the Slack API Website**: Go to the [Slack API Apps Page](https://api.slack.com/apps) and sign in with your Slack account.
2. **Create a New App**:
   - Click on **"Create New App"**.
   - Choose **"From scratch"**.
   - Enter an **App Name** (e.g., `Chatz Bot`).
   - Choose the **Slack Workspace** where the app will be used.
   - Click **"Create App"**.

## Step 2: Configure OAuth & Permissions

1. **Navigate to OAuth & Permissions**:
   - In the left sidebar, click **"OAuth & Permissions"**.
2. **Add Bot Token Scopes**:
   - Scroll down to the **Scopes** section.
   - Under **Bot Token Scopes**, click **"Add an OAuth Scope"** and add the following scopes based on your requirements:
     - `chat:write`: Allows the bot to send messages to channels.
     - `channels:read`: Allows the bot to view channel information (optional).
     - `groups:read`: Allows the bot to view private group information (optional).
     - `chat:write.customize`: Allows the bot to send messages with customized formatting (optional).
3. **Install the App to Your Workspace**:
   - Scroll up to **OAuth Tokens for Your Workspace**.
   - Click **"Install App to Workspace"**.
   - Review the permissions and click **"Allow"** to grant your app access.

## Step 3: Obtain the Bot Token

1. **Copy the Bot Token**:
   - After installation, you will see the **Bot User OAuth Token** under the **OAuth Tokens for Your Workspace** section.
   - Copy this token. It will look something like `xoxb-XXXXXXXXXXXX-XXXXXXXXXXXX-XXXXXXXXXXXXXX`.

## Step 4: Configure the Bot in Slack

1. **Navigate to App Home**:
   - Go to the **App Home** tab in your Slack App configuration page.
   - Click on **"Review Scopes to Add Bot"** to ensure your bot has the correct scopes.
2. **Customize Your Bot**:
   - Give your bot a username, and customize its icon if desired.

## Step 5: Adding the Bot to Channels

1. **Manually Add the Bot**:
   - To start interacting with the bot, invite it to specific channels using the Slack command:
     ```
     /invite @YourBotName
     ```

## Step 6: Using the Bot Token in Your Application (e.g., Chatz)

1. **Set Up Environment Variables**:
   - Use the Bot Token you copied earlier in your application by setting it as an environment variable. For example, add the following line to your `~/.chatz.ini` file:
     ```ini
     [default]
     PROVIDER=slack
     TOKEN=xoxb-XXXXXXXXXXXX-XXXXXXXXXXXX-XXXXXXXXXXXXXX
     CHANNEL_ID=XXXXXXXXX
     ```
2. **Example Configuration in Chatz**:
   - If using Chatz, ensure the application reads the token correctly from the environment variables and is configured to send messages to the right Slack channels.

## Step 7: Testing Your Slack Bot

1. **Send a Test Message**:
    ```bash
    chatz -o "Test message"
    ```

## Troubleshooting Tips

- **Invalid Token Error**: Ensure that the bot token is correctly set and has the required permissions.
- **Bot Not Responding**: Verify that the bot is invited to the correct channels and that the Slack app permissions are set correctly.
- **Scope Issues**: Review the scopes added in the OAuth & Permissions section to ensure the bot has adequate access.

## Security Considerations

- Keep your Slack Bot Token secret. Never expose it in public repositories or share it openly.
- Regenerate the token if you suspect it has been compromised.
