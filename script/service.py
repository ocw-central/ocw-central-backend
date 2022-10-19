import google.auth
from googleapiclient.discovery import build


def get_service():
    credentials, _ = google.auth.default()
    return build('drive', 'v3', credentials=credentials)
