# covid-skill

COVID-19 search skill

## Deploy Cloud Function

```shell script
gcloud functions deploy covid-skill --runtime go113 --trigger-http --region europe-west3 --memory 128MB --entry-point Cases
```
