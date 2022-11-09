import logging

from mysql.connector import connect


def main():
    connection = connect(
        host="127.0.0.1",
        user="root",
        password="",
        port=3306,
        database="ocw-central",
    )
    ids = set()
    select_resource_query = "select id from subjects limit 1"
    with connection.cursor() as cursor:
        cursor.execute(select_resource_query)
        res = list(cursor.fetchall())

        while len(res) > 0:
            id = bytes(res.pop()[0]).hex()
            if id in ids:
                continue

            ids.add(id)
            sql = (
                "select related_subject_id "
                "from subject_related_subjects "
                f"where subject_id = 0x{id}"
            )
            cursor.execute(sql)
            res.extend(list(cursor.fetchall()))

    print(len(ids))
    for id in ids:
        print(id)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    main()
