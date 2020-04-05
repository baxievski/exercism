from datetime import date, timedelta


class MeetupDayException(Exception):
    pass


def meetup(year, month, week, day_of_the_week):
    day_map = {
        "Monday": 0,
        "Tuesday": 1,
        "Wednesday": 2,
        "Thursday": 3,
        "Friday": 4,
        "Saturday": 5,
        "Sunday": 6,
    }
    week_map = {
        "1st": 0,
        "2nd": 1,
        "3rd": 2,
        "4th": 3,
        "5th": 4,
        "last": -1,
    }
    days = [
        day
        for day in (date(year, month, 1) + timedelta(days=x) for x in range(32))
        if day.weekday() == day_map[day_of_the_week] and day.month == month
    ]
    if week == "teenth":
        for day in days:
            if 12 < day.day < 20:
                return day
    if len(days) <= week_map[week]:
        raise MeetupDayException("Nonexistant day")
    return days[week_map[week]]


if __name__ == "__main__":
    print(f'{meetup(2015, 2, "5th", "Monday")=}')
