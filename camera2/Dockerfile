
MAINTAINER camlistore <camlistore@googlegroups.com>

CMD ["./gradlew"]

RUN echo "Adding gopher user and group" \
	&& groupadd --system --gid 1000 gopher \
	&& useradd --system --gid gopher --uid 1000 --shell /bin/bash --create-home gopher \
	&& mkdir /home/gopher/.gradle \
	&& chown --recursive gopher:gopher /home/gopher

# To enable running android tools such as aapt
RUN apt-get update && apt-get -y upgrade
RUN apt-get install -y lib32z1 lib32stdc++6
# For Go:
RUN apt-get -y --no-install-recommends install curl gcc
RUN apt-get -y --no-install-recommends install ca-certificates libc6-dev git

USER gopher
VOLUME "/home/gopher/.gradle"
ENV GOPHER /home/gopher

# Get android sdk, ndk, and rest of the stuff needed to build the android app.
WORKDIR $GOPHER
RUN mkdir android-sdk
ENV ANDROID_HOME $GOPHER/android-sdk
#RUN echo y | $ANDROID_HOME/tools/bin/sdkmanager 'build-tools;22.0.1'
RUN echo y | $ANDROID_HOME/tools/bin/sdkmanager 'build-tools;23.0.3'
RUN echo y | $ANDROID_HOME/tools/bin/sdkmanager 'extras;android;m2repository'
RUN echo y | $ANDROID_HOME/tools/bin/sdkmanager 'ndk-bundle'

# Get gradle. We don't actually need to build the app, but we need it to
# generate the gradle wrapper, since it's not included in the app's repo.
WORKDIR $GOPHER
ENV GRADLE_VERSION 4.1
ARG GRADLE_DOWNLOAD_SHA256=d55dfa9cfb5a3da86a1c9e75bb0b9507f9a8c8c100793ccec7beb6e259f9ed43
RUN set -o errexit -o nounset \
	&& echo "Downloading Gradle" \
	&& wget --no-verbose --output-document=gradle.zip "https://services.gradle.org/distributions/gradle-${GRADLE_VERSION}-bin.zip"
RUN echo "Checking download hash" \
	&& echo "${GRADLE_DOWNLOAD_SHA256} *gradle.zip" | sha256sum --check -
RUN echo "Installing Gradle" \
	&& unzip gradle.zip \
	&& rm gradle.zip
RUN mkdir $GOPHER/bin \
	&& ln --symbolic "${GOPHER}/gradle-${GRADLE_VERSION}/bin/gradle" $GOPHER/bin/gradle

# Get Go stable release
WORKDIR $GOPHER
RUN curl -O https://storage.googleapis.com/golang/go1.11.1.linux-amd64.tar.gz
RUN echo '2871270d8ff0c8c69f161aaae42f9f28739855ff5c5204752a8d92a1c9f63993  go1.11.1.linux-amd64.tar.gz' | sha256sum -c
RUN tar -xzf go1.11.1.linux-amd64.tar.gz
ENV GOPATH $GOPHER
ENV GOROOT $GOPHER/go
ENV PATH $PATH:$GOROOT/bin:$GOPHER/bin

# Get gomobile
RUN go get -u golang.org/x/mobile/cmd/gomobile
WORKDIR $GOPATH/src/golang.org/x/mobile/cmd/gomobile
RUN git reset --hard 069be623eb8e75049d64f1419849b3e92aab1c81
RUN go install

# init gomobile
RUN gomobile init -ndk $ANDROID_HOME/ndk-bundle


