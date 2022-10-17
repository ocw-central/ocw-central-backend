import json
from pathlib import Path

from mysql.connector import connect

from upload_file import upload_file_from_link


def main():
    connection = connect(
        host="127.0.0.1",
        user="root",
        password="",
        port=3307,
        database="ocw-central",
    )
    id_links = []
    select_resource_query = "select id, title, link from resources"
    with connection.cursor() as cursor:
        cursor.execute(select_resource_query)
        result = cursor.fetchall()
        for (id, title, link) in result:
            file_name = Path(link).name
            shared_link = upload_file_from_link(link, file_name)
            id_links.append({
                "id": id,
                "title": title,
                "file_name": file_name,
                "link": shared_link,
            })

    connection.close()
    with open("id_links.json", "w") as f:
        json.dump(id_links, f)


if __name__ == "__main__":
    main()
