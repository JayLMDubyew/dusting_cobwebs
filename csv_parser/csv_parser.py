import csv

def main():
    with open('data-messy.csv',r) as infile:
        csvreader = csv.reader(infile,dialect="excel")