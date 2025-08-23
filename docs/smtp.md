# SMTP with Gmail

This document explains how to configure SMTP with Gmail, including creating an App Password, and understanding TLS/STARTTLS and encryption options.

## 1. Creating an App Password

Gmail requires an "App Password" for less secure apps accessing your account. This is a more secure alternative to using your regular Gmail password directly.

1. Go to your [Google account security settings](https://myaccount.google.com/security).
2. Scroll down to "Signing in to Google" and click "App Passwords".
3. Select "Mail" as the app and "Other (Custom name)" as the device. Give it a descriptive name (e.g., "My SMTP App").
4. Click "Generate". A new password will be displayed. **Copy this password immediately; you won't be able to see it again.**

## 2. SMTP Server Settings

- **Server:** `smtp.gmail.com`
- **Port:**
  - **STARTTLS:** 587 (Recommended)
  - **TLS:** 465
  - **No Encryption (Insecure - Avoid):** 25
- **Username:** Your full Gmail address (e.g., `yourname@gmail.com`)
- **Password:** The App Password you generated.

## 3. TLS/STARTTLS and Encryption

- **TLS (Transport Layer Security) / STARTTLS:** These are encryption protocols that secure the connection between your email client and the SMTP server. They encrypt your email messages and prevent eavesdropping. **Using TLS/STARTTLS is strongly recommended.** Many email clients support STARTTLS, initiating encryption during the connection process.

- **No Encryption:** Sending emails without encryption is highly discouraged. Your email, including the subject, body, and any attachments, could be intercepted by malicious actors. Avoid this unless absolutely necessary, and only with trusted parties.

## 4. Example Configuration (Illustrative - adapt to your email client)

This is a generic example; the exact settings will depend on your email client (e.g., Outlook, Thunderbird). Refer to your email client's documentation for specific instructions.

```
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USE_STARTTLS=true
```

Remember to replace placeholders with your actual information.

### Using CLI Flags

You can also override the subject using the CLI flag:

```sh
chatz --profile=smtp --subject "Custom Subject" "Hello from Chatz!"
```

This will send an email with "Custom Subject" as the subject.

## 5. Troubleshooting

If you encounter issues, check:

- **Correct App Password:** Verify you copied the correct App Password.
- **Firewall Settings:** Ensure your firewall isn't blocking outgoing connections on port 587 (or 465 if using no encryption).
- **Email Client Configuration:** Double-check all server settings in your email client.

Using an App Password is crucial for securing your Gmail account when using SMTP. Always prioritize using TLS/STARTTLS encryption for the security of your emails.

```

```
