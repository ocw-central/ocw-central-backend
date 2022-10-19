from typing import List

from service import get_service

PAGE_SIZE = 1000


def get_file_ids() -> List[str]:
    service = get_service()
    file_ids = []
    page_token = None
    while True:
        results = service.files() \
            .list(
                pageSize=PAGE_SIZE,
                pageToken=page_token,
                fields="nextPageToken, files(id, name)"
            ) \
            .execute()

        assert PAGE_SIZE == len(results["files"])
        file_ids.extend([file["id"] for file in results["files"]])

        page_token = results.get("nextPageToken", None)
        if page_token is None:
            break

    return file_ids
