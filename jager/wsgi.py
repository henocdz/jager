"""
WSGI config for jager project.

It exposes the WSGI callable as a module-level variable named ``application``.

For more information on this file, see
https://docs.djangoproject.com/en/2.0/howto/deployment/wsgi/
"""

import os

from django.core.wsgi import get_wsgi_application

ENVS = {
    'production': 'jager.settings.production',
    'staging': 'jager.settings.staging',
    'development': 'jager.settings.development',
}

env = os.getenv('ENVIRONMENT', 'development')
settings_module = ENVS[env]
os.environ.setdefault('DJANGO_SETTINGS_MODULE', settings_module)

application = get_wsgi_application()
