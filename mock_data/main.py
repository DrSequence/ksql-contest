import uuid

def generate_uuids():
    return [str(uuid.uuid4()) for _ in range(100)]

def save_uuids_to_file(filename, uuids):
    with open(filename, 'w') as file:
        file.write(','.join(uuids))

def main():
    uuids = generate_uuids()
    save_uuids_to_file('uuids.csv', uuids)
    print(f"UUIDs have been saved to 'uuids.csv'.")

if __name__ == '__main__':
    main()
