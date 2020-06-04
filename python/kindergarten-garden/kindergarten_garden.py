class Garden(object):
    plants_map = {
        "G": "Grass",
        "C": "Clover",
        "R": "Radishes",
        "V": "Violets",
    }

    def __init__(self, diagram, students=None):
        if students is None:
            students = [
                "Alice",
                "Bob",
                "Charlie",
                "David",
                "Eve",
                "Fred",
                "Ginny",
                "Harriet",
                "Ileana",
                "Joseph",
                "Kincaid",
                "Larry",
            ]
        self.students = sorted(students)
        self.diagram = diagram.split("\n")

    def plants(self, student):
        i = self.students.index(student)
        first_row_plants = self.diagram[0][2 * i : 2 * i + 2]
        second_row_plants = self.diagram[1][2 * i : 2 * i + 2]
        student_plants = "{}{}".format(first_row_plants, second_row_plants)

        return [Garden.plants_map[x] for x in student_plants]


if __name__ == "__main__":
    garden = Garden("RC\nGG")
    print(f"{garden=}")
    print(f'{garden.plants("Alice")=}')
