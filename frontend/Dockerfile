FROM google/dart AS build-env

RUN apt-get update && \
    apt-get install -y curl git wget zip unzip libgconf-2-4 gdb libstdc++6 libglu1-mesa fonts-droid-fallback \
        lib32stdc++6 python3 && \
    apt-get clean

RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter

RUN /usr/local/flutter/bin/flutter doctor -v
ENV PATH="/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin:${PATH}"

RUN flutter channel master
RUN flutter upgrade
RUN flutter config --enable-web

WORKDIR /usr/local/influ_dojo
COPY . /usr/local/influ_dojo
RUN flutter build web

FROM nginx
COPY --from=build-env /usr/local/influ_dojo/build/web /usr/share/nginx/html
