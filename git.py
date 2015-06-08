
import os
import subprocess
from contextlib import ContextDecorator


class cd(ContextDecorator):
    """Context manager/decorator for changing the current working directory."""
    def __init__(self, new_path):
        self.new_path = os.path.expanduser(new_path)

    def __enter__(self):
        self.previous_path = os.getcwd()
        os.chdir(self.new_path)
        return self

    def __exit__(self, *exc):
        os.chdir(self.previous_path)
        return False


# TODO use same context for all methods
class git:
    """Minimal git wrapper, providing only funtions to init add and commit."""
    path = ''
    def __init__(self, path):
        assert os.path.isabs(path)
        self.path = path

    @cd(path)
    def init(self):
        print('current dir: {}'.format(os.getcwd()))
        subprocess.call('git init')


    def add(self, filenames):
        for filename in filenames:
            assert os.path.isfile(filename)  # needs entire path
            subprocess.call('git add ' + filename)

    def commit(self, date, message=''):
        subprocess.call(
            "git commit -m '{m}' --date {d}".format(m=message, d=date)
        )
