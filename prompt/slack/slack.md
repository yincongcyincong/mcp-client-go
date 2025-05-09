1. **`slack_list_channels`** – List public or pre-defined channels in the workspace with pagination.

    - List channels in the workspace with a limit of 50 channels, using pagination if necessary.

2. **`slack_post_message`** – Post a new message to a Slack channel.

    - Post the message "Hello, everyone!" to the channel with ID `C12345678`.

3. **`slack_reply_to_thread`** – Reply to a specific message thread in Slack.

    - Reply to the thread in the channel `C12345678` with the message "Thanks for the update!" on the message with timestamp `1234567890.123456`.

4. **`slack_add_reaction`** – Add a reaction emoji to a message.

    - Add the emoji `thumbsup` reaction to the message with timestamp `1234567890.123456` in the channel `C12345678`.

5. **`slack_get_channel_history`** – Get recent messages from a channel.

    - Get the last 20 messages from the channel with ID `C12345678`.

6. **`slack_get_thread_replies`** – Get all replies in a message thread.

    - Retrieve all replies in the thread with timestamp `1234567890.123456` in the channel `C12345678`.

7. **`slack_get_users`** – Get a list of all users in the workspace with their basic profile information.

    - Get a list of up to 100 users in the workspace, with pagination if necessary.

8. **`slack_get_user_profile`** – Get detailed profile information for a specific user.

    - Retrieve the profile information for the user with ID `U12345678`.
