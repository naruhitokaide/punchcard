
import os
import re
import subprocess


DATE_REGEX = '\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}'  # ISO 8601


class cd:
    """Context manager for changing the current working directory."""
    def __init__(self, new_path):
        self.new_path = os.path.expanduser(new_path)

    def __enter__(self):
        self.previous_path = os.getcwd()
        os.chdir(self.new_path)

    def __exit__(self, etype, value, traceback):
        os.chdir(self.previous_path)


class git:
    """Minimal git wrapper, providing only: init, add and commit."""

    def __init__(self, path):
        # TODO create path if it does not exist?
        assert os.path.isabs(path)
        self.path = path

    def init(self):
        with cd(self.path):
            subprocess.call('git init')

    def add(self, filenames):
        with cd(self.path):
            for filename in filenames:
                assert os.path.isfile(filename)  # needs entire path
                subprocess.call('git add ' + filename)

    def commit(self, date, message=''):
        assert re.compile(DATE_REGEX).match(date)
        with cd(self.path):
            subprocess.call(
                "git commit -m '{m}' --date {d}".format(m=message, d=date)
            )
