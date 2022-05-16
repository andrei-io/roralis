# Run from top level folder in monorepo
# TODO script to automatically extract from downloaded keystore
KEY_ALIAS=`cat ./keystore/password.alias`
rm android-out/bundle.apks
java -jar "../../../Programs/android-stuff/bundletool.jar" build-apks --bundle=./android-out/roralis-0.5.aab --output=./android-out/bundle.apks --mode=universal --ks=./keystore/keys.jks --ks-pass=file:./keystore/keystore.pwd --ks-key-alias=$KEY_ALIAS --key-pass=file:./keystore/key.pwd
unzip -p ./android-out/bundle.apks universal.apk > ./android-out/Roralis.apk