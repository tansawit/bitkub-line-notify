# BitKub Line Notify Service

A service to send daily notification on the value of your BitKub cryptocurrency wallet.

## Usage

Clone the repositoy:

```bash
git clone https://github.com/tansawit/bitkub-line-notify
```

Create a `.env` file inside the cloned `bitkub-line-notify` folder with the following information after registering with the BitKub Public API [page](https://www.bitkub.com/publicapi)

- `BITKUB_API_KEY`
- `BITKUB_API_SECRET`

Create a [new LINE notify service](https://notify-bot.line.me/my/) by generating a new access token. Add that token to the `.env` file as `LINE_NOTIFY_TOKEN`

Then you can run the Dockerfile on a machine/cloud service of your choice and it will send you details about your portfolio every day at 9:00 AM
