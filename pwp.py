#!/usr/bin/python

import os
import hashlib
import getpass

password_length = 12 - 5

base_phrase = getpass.getpass( "Enter base phrase: " )

sanity_check = hashlib.sha256()
sanity_check.update( base_phrase )

print sum( [ ord( x ) for x in sanity_check.digest() ] )

valid_characters = 'abcdefghijklmnopqrstuvwxyz0123456789'

while ( True ) :
  door_id = raw_input( "Enter door id: " )
  if ( door_id == '' ) :
    break

  key_data = hashlib.sha256()
  key_data.update( base_phrase + ' - ' + door_id )

  print ''.join( [ valid_characters[ ord( x ) % len( valid_characters ) ] for x in key_data.digest() ][:password_length] )

os.system( 'clear' )

