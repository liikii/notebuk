"""
argument type check.
"""


class ArgumentTypeError(BaseException):
    """ Inappropriate argument value (of correct type). """
    def __init__(self, *args, **kwargs):  # real signature unknown
        pass


def check_type(*tps):
    def real_decorator(y):
        def wrap_func(*args):
            if len(tps) != len(args):
                raise Exception('check type list length not matched.')
            for i in range(len(args)):
                if not isinstance(args[i], tps[i]):
                    raise ArgumentTypeError('argument type error. argument %s need type %s ' % (i, tps[i]))
            return y(*args)
        return wrap_func
    return real_decorator


@check_type(int, int)
def add_int(a, b):
    return a+b


@check_type(int, str)
def bad_add_int(a, b):
    return a+b


print(add_int(3, 4))
print(bad_add_int(3, 4))
