import json
import logging
from pathlib import Path

from mysql.connector import connect

from upload_file import upload_file_from_link


def main():
    connection = connect(
        host="127.0.0.1",
        user="root",
        password="",
        port=3306,
        database="ocw-central",
    )
    id_links = []
    select_resource_query = "select id, title, link from resources"
    with connection.cursor() as cursor:
        cursor.execute(select_resource_query)
        result = cursor.fetchall()
        for (resource_id, title, link) in result:
            file_name = Path(link).name
            shared_link, file_id = upload_file_from_link(link, file_name)

            logging.info(
                f"Uploaded {file_name} with id:{file_id} and link:{shared_link}."
            )

            id_links.append({
                "resource_id": resource_id,
                "drive_file_id": file_id,
                "title": title,
                "file_name": file_name,
                "link": shared_link,
            })

    with connection.cursor() as cursor:
        for id_link in id_links:
            update_resource_query = (
                "update resources "
                f"set link = '{id_link['link']}' "
                f"where id = {id_link['resource_id']}"
            )
            cursor.execute(update_resource_query)

        connection.commit()

    connection.close()
    with open("id_links.json", "w") as f:
        json.dump(id_links, f)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    main()
