#!/usr/bin/env python
import re
import os
import sys


def load_env():
    try:
        with open('.env') as env_file:
            env_lines = env_file.read()
    except IOError:
        env_lines = ''

    for env_var in env_lines.splitlines():
        line = re.match(r'\A([A-Za-z_0-9]+)=(.*)\Z', env_var)
        if line:
            key, value = line.group(1), line.group(2)
            if value.startswith("'"):
                value = value[1:-1]
            os.environ.setdefault(key, value)
        elif env_var:
            print('Invalid env var: {}'.format(env_var))



if __name__ == "__main__":
    os.environ.setdefault("DJANGO_SETTINGS_MODULE", "jager.settings.development")
    try:
        from django.core.management import execute_from_command_line
    except ImportError as exc:
        raise ImportError(
            "Couldn't import Django. Are you sure it's installed and "
            "available on your PYTHONPATH environment variable? Did you "
            "forget to activate a virtual environment?"
        ) from exc


    load_env()
    execute_from_command_line(sys.argv)
