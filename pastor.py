#!/usr/bin/python

import os
import hashlib
import getpass
import pyperclip
from pbkdf2 import crypt

password_length = 16

base_phrase = getpass.getpass("Enter base phrase: ")

sanity_check = hashlib.sha256()
sanity_check.update(base_phrase)

print sum([ord(x) for x in sanity_check.digest()])

valid_characters = 'abcdefghijklmnopqrstuvwxyz0123456789.~!@#$%^&*()_+'

while (True):
    door_id = raw_input("Enter door id: ")
    if (door_id == ''):
        break

    key = crypt(door_id, base_phrase, 1000)
    key_data = hashlib.sha256()
    key_data.update(key)

    password = ''.join([valid_characters[ord(x) % len(valid_characters)] for x in key_data.digest()][:password_length])
    pyperclip.setcb(password)
    print pyperclip.getcb()

os.system('clear')
