"""
desc: args type checker.
"""


class ArgumentTypeError(BaseException):
    def __init__(self, *args, **kwargs):
        pass


class CheckTypeLengthError(BaseException):
    def __init__(self, *args, **kwargs):
        pass


def check_type(*tps):
    def real_decorator(y):
        def wrap_func(*args):
            if len(tps) != len(args):
                raise CheckTypeLengthError('check type list length not matched.')
            for i in range(len(args)):
                if not isinstance(args[i], tps[i]):
                    raise ArgumentTypeError('argument type error. argument %s need type %s ' % (i, tps[i]))
            return y(*args)
        return wrap_func
    return real_decorator


@check_type(int, int)
def add_int(a, b):
    return a+b


@check_type(str, str)
def add_str(a, b):
    return a + ' + ' + b


# print(add_int(3, 'aaa'))
print(add_str('3', '4'))
