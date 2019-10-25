# Serverless video downloader (`youtube-dl`)

A serverless application that uses `youtube-dl` to quickly download and stream
videos from media sites.

Deploy to Google Cloud Run:

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)

Then visit your application with `?url=` parameter to download a video:

    https://<YOUR-APP-URL>/?url=https://www.youtube.com/watch?v=jHjFxJVeCQs

Currently the maximum video length is set to 60 seconds because of
32 MB [response size limit](https://cloud.google.com/run/quotas) on Cloud Run.

---

This project is not affiliated with Google. It's created
for demonstration purposes.
