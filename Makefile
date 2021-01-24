all: install

install: script/install-sm.rb
		@ruby script/install-sm.rb
