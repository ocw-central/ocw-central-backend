import google.auth
from googleapiclient.discovery import build


def get_service():
    credentials, _ = google.auth.default()
    return build('drive', 'v3', credentials=credentials)


def get_about():
    """
    https://developers.google.com/drive/api/v3/reference/about
    """
    return get_service().about().get(fields="storageQuota").execute()


if __name__ == "__main__":
    print(get_about())
