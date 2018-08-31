"""
function exercises.
"""


def g(a, b):
    # wrapper
    def wf(f):
        # replace
        def rf(*args):
            print('len', len(args))
            p = args[0]
            if a < p < b:
                print('a < p < b')
            else:
                print('out interval')
            f(p)
        return rf
    return wf


@g(3, 9)
def ha_ha(t=0):
    print(t)


ha_ha(5)
ha_ha(9)
