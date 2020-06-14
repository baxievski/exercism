class Table:
    def __init__(self):
        self._teams = {}

    def add_match(self, t1_name, t2_name, result):
        t1 = Team(t1_name)
        if t1_name in self._teams:
            t1 = self._teams[t1_name]

        t2 = Team(t2_name)
        if t2_name in self._teams:
            t2 = self._teams[t2_name]

        if result == "draw":
            t1.add_draw()
            t2.add_draw()
        elif result == "win":
            t1.add_win()
            t2.add_loss()
        elif result == "loss":
            t2.add_win()
            t1.add_loss()
        else:
            raise Exception(f"unrecognized result: {result}")

        self._teams[t1_name], self._teams[t2_name] = t1, t2

    def render(self):
        header = f"{'Team': <30}"
        header += f" | {'MP': >2}"
        header += f" | {'W': >2}"
        header += f" | {'D': >2}"
        header += f" | {'L': >2}"
        header += f" | {'P': >2}"

        results = [
            header,
        ]
        for team in sorted(self._teams.values(), reverse=True):
            row = f"{str(team): <30}"
            row += f" | {team.matches_played: >2}"
            row += f" | {team.wins: >2}"
            row += f" | {team.draws: >2}"
            row += f" | {team.losses: >2}"
            row += f" | {team.points: >2}"
            results.append(row)

        return results


class Team:
    results_to_points = {"win": 3, "draw": 1, "loss": 0}

    def __init__(self, name):
        self.name = name
        self.matches_played = 0
        self.wins = 0
        self.losses = 0
        self.draws = 0

    @property
    def points(self):
        pts = Team.results_to_points["win"] * self.wins
        pts += Team.results_to_points["draw"] * self.draws
        pts += Team.results_to_points["loss"] * self.losses

        return pts

    def add_win(self):
        self.matches_played += 1
        self.wins += 1

    def add_draw(self):
        self.matches_played += 1
        self.draws += 1

    def add_loss(self):
        self.matches_played += 1
        self.losses += 1

    def __repr__(self):
        return str(self.name)

    def __hash__(self):
        return hash(repr(self))

    def __lt__(self, other):
        a = (-1 * self.points, self.matches_played, self.name)
        b = (-1 * other.points, other.matches_played, other.name)

        return b < a


def tally(matches):
    table = Table()
    for match in matches:
        table.add_match(*match.split(";"))

    return table.render()


def main():
    table = tally(
        [
            "Courageous Californians;Devastating Donkeys;win",
            "Allegoric Alaskans;Blithering Badgers;win",
            "Devastating Donkeys;Allegoric Alaskans;loss",
            "Courageous Californians;Blithering Badgers;win",
            "Blithering Badgers;Devastating Donkeys;draw",
            "Allegoric Alaskans;Courageous Californians;draw",
        ]
    )
    for row in table:
        print(row)


if __name__ == "__main__":
    main()
