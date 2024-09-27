from faker import Faker
from feedgen.feed import FeedGenerator
import random

fake = Faker()

def generate_random_feeds(fg) -> str:
    for _ in range(5):
        fe = fg.add_entry()
        fe.title(fake.sentence(nb_words=10))
        fe.link(href=fake.url())
        fe.description(fake.paragraph())

def main():
    fg = FeedGenerator()
    fg.title(fake.sentence(nb_words=10))
    fg.link(href=fake.url())
    fg.description(fake.paragraph())

    for _ in range(10):
        generate_random_feeds(fg)

    fg.rss_file(filename="rss_feed.xml",pretty=True)

if __name__ == "__main__":
    main()