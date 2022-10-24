import json
import logging
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

        logging.info(f"Fetched {len(results['files'])} files.")
        file_ids.extend([file["id"] for file in results["files"]])

        page_token = results.get("nextPageToken", None)
        if page_token is None:
            break

    logging.info(f"Finished fetching. Fetched {len(file_ids)} files.")
    return file_ids


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    ids = get_file_ids()
    with open("uploaded_file_ids.json", "w") as f:
        json.dump(ids, f)
