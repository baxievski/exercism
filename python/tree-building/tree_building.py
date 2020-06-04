class Record:
    def __init__(self, record_id, parent_id):
        self.record_id = record_id
        self.parent_id = parent_id

    def validate(self):
        if self.record_id == self.parent_id != 0:
            raise ValueError("Records can't be their own parrents")

        if self.record_id < self.parent_id:
            raise ValueError("Records must have higher ids than their parrents'")

    def __str__(self):
        return f"{self.record_id}, {self.parent_id}"

    def __eq__(self, other):
        if self.record_id != other.record_id:
            return False

        if self.parent_id != other.parent_id:
            return False

        return True

    def __lt__(self, other):
        a = (self.record_id, self.parent_id)
        b = (other.record_id, other.parent_id)

        return a < b

    def __le__(self, other):
        a = (self.record_id, self.parent_id)
        b = (other.record_id, other.parent_id)

        return a <= b

    def __gt__(self, other):
        a = (self.record_id, self.parent_id)
        b = (other.record_id, other.parent_id)

        return a > b

    def __ge__(self, other):
        a = (self.record_id, self.parent_id)
        b = (other.record_id, other.parent_id)

        return a >= b


class Node:
    def __init__(self, node_id):
        self.node_id = node_id
        self.children = []


def BuildTree(records):
    nodes = {}
    root = None

    for i, record in enumerate(sorted(records)):
        record.validate()
        nodes[record.record_id] = Node(record.record_id)
        if i != record.record_id:
            raise ValueError("Node is out of order")

        if record.record_id == 0:
            root = nodes[0]
            continue

        parrent = nodes[record.parent_id]
        parrent.children.append(nodes[record.record_id])

    return root


if __name__ == "__main__":
    pass
