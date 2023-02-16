FROM ruby:2.7.5

RUN apt-get update && apt-get install -y nodejs

WORKDIR /app

COPY Gemfile* ./

RUN gem install bundler -v 2.1.4

RUN bundle install

COPY . .

EXPOSE 3000

CMD ["bundle", "exec", "rails", "s", "-b", "0.0.0.0"]
