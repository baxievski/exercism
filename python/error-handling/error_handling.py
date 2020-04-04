def handle_error_by_throwing_exception():
    raise Exception("Error")


def handle_error_by_returning_none(input_data):
    if input_data in "0123456789":
        return int(input_data)
    return None


def handle_error_by_returning_tuple(input_data):
    if input_data in "0123456789":
        return (True, int(input_data))
    return (False, None)


def filelike_objects_are_closed_on_exception(filelike_object):
    with filelike_object as f:
        f.do_something()


if __name__ == "__main__":
    print(f"main()")
