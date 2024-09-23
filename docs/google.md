# Creating and Configuring a Google Chat Webhook

Google Chat webhooks allow you to send messages directly to Google Chat spaces (rooms or direct messages) from external applications. This guide will help you set up a Google Chat webhook URL and configure it for use in your project.

## Step 1: Set Up a Google Chat Space

1. **Open Google Chat**: Go to [Google Chat](https://chat.google.com/) and sign in with your Google account.
2. **Create a New Space** (if you don't have one):
   - Click on the **"+"** button next to "Spaces" in the left sidebar.
   - Select **"Create Space"**.
   - Enter a **Space Name** (e.g., `Chatz Notifications`).
   - Add members if needed (optional).
   - Click **"Create"**.

## Step 2: Configure the Webhook

1. **Open Space Settings**:
   - Once the space is created, click on the space name at the top of the chat window.
   - Click on **"Manage Webhooks"**. If you do not see this option, make sure you have the necessary permissions or try creating a new space.

2. **Create a New Webhook**:
   - Click on **"Configure Webhook"**.
   - Enter a **Webhook Name** (e.g., `Chatz Webhook`).
   - Optionally, upload an image URL to customize the webhook avatar.

3. **Generate the Webhook URL**:
   - After entering the details, click **"Save"**.
   - Copy the generated **Webhook URL**. It will look something like this: 
     ```
     https://chat.googleapis.com/v1/spaces/XXXXXX/messages?key=XXXXXXXX&token=XXXXXXXX
     ```

## Step 3: Using the Webhook URL in Your Application (e.g., Chatz)

1. **Set Up Environment Variables**:
   - Store the Webhook URL in an environment variable for easy access. For example, add the following line to your `~/.chatz.ini` file:
     ```ini
     [default]
     PROVIDER=google
     WEB_HOOK_URL=https://chat.googleapis.com/v1/spaces/XXXXXX/messages?key=XXXXXXXX&token=XXXXXXXX
     ```

2. **Example Configuration in Chatz**:
   - Ensure your application reads the webhook URL correctly from the environment variables and is configured to send messages to Google Chat.

## Step 4: Testing Your Google Chat Webhook

1. **Send a Test Message**:
    ```bash
    chatz -o "Test message."
    ```

## Troubleshooting Tips

- **Invalid Webhook URL**: Double-check the URL for any errors, such as missing characters or incorrect tokens.
- **Insufficient Permissions**: Ensure you have the correct permissions to create and manage webhooks in your Google Chat space.
- **Message Not Sent**: Verify that your payload is correctly formatted as JSON and that you are using the correct headers.

## Security Considerations

- Keep your Google Chat webhook URL secret. Never expose it in public repositories or share it openly.
- If the webhook URL is compromised, delete it and generate a new one.
