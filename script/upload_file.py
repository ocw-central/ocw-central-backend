import requests
import tempfile

import google.auth
from googleapiclient.discovery import build
from googleapiclient.http import MediaFileUpload


def get_service():
    credentials, _ = google.auth.default()
    return build('drive', 'v3', credentials=credentials)


def upload_file_from_link(link: str, file_name: str) -> str:
    response = requests.get(link)
    fp = tempfile.NamedTemporaryFile()
    fp.write(response.content)
    path = fp.name
    shared_link = upload_file(path, file_name)
    fp.close()
    return shared_link


def upload_file(path: str, file_name: str) -> str:
    service = get_service()
    file_metadata = {'name': file_name}
    media = MediaFileUpload(path)
    file = service.files() \
        .create(body=file_metadata, media_body=media, fields='id') \
        .execute()

    file = service.permissions() \
        .create(fileId=file["id"], body={'type': 'anyone', 'role': 'reader'}) \
        .execute()

    return file["webViewLink"]
