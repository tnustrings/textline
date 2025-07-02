name = "lily"

# get the version from github tag
# delete the v from the version tag cause python build seems to strip it as well
version = $(shell git tag | tail -1)


all: lily.deb

lily.deb: ../bin/lily
	# create a directory for the package and in it a DEBIAN directory with meta info.
	mkdir -p "lily-${version}"
	mkdir -p "lily-${version}/DEBIAN"

	# make the bin

	mkdir -p "lily-${version}/usr/local/bin"

	# put the lily executable in

	cp ../bin/lily "lily-${version}/usr/local/bin"

	# manually create a file lily-<version>/deb/DEBIAN/control and put in the package info.
	# build the package
	dpkg-deb --build "lily-${version}"


release: 
	make
	# push the tag
	git push origin tag "${version}"
	gh release create "${version}" "lily-${version}.deb"

release-update: # if an asset was already uploaded, delete it before uploading again
	gh release delete-asset "${version}" "lily-${version}.deb" -y
	make release


