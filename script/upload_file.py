import tempfile
from typing import List

import requests

from googleapiclient.http import MediaFileUpload

from service import get_service


def upload_file_from_links(links: List[str], file_name: str) -> str:
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
