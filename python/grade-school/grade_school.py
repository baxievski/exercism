class School(object):
    def __init__(self):
        self.students = {
            1: set(),
            2: set(),
            3: set(),
            4: set(),
            5: set(),
            6: set(),
            7: set(),
            8: set(),
            9: set(),
        }

    def add_student(self, name, grade):
        self.students[grade].add(name)

    def grade(self, grade):
        return sorted(self.students[grade])

    def roster(self):
        all_grades = []
        for _, students in self.students.items():
            if not students:
                continue
            all_grades.extend(sorted(students))
        return all_grades


if __name__ == "__main__":
    school = School()
    school.add_student(name="Bojan", grade=1)
    school.add_student(name="Bojan", grade=2)
    school.add_student(name="Sime", grade=2)
    print(f"{school.students=}")
    print(f"{school.roster()=}")
