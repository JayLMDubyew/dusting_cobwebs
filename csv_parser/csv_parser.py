import csv
import re
from dataclasses import dataclass

@dataclass
class Employee:
    name:str
    species:str
    day:int
    status:str


def main():
    personnel = []
    with open('data-messy.csv','r') as infile:
        csvreader = csv.reader(infile,dialect="excel")
        _ = next(csvreader)

        for row in csvreader:
            
            if len(row) < 4:
                print(f'Row is malformed - too short: {row}')
                continue
            else:
                if re.fullmatch(r'[0-9]+$',row[2].strip()) is None:
                    print(f'Improper or non-positive value given for day: {row[2]}')
                newEmployee = Employee(row[0].strip(),row[1].strip(),row[2].strip(),row[3].strip())
                personnel.append(newEmployee)    
    print(personnel)
    count_status(personnel)                

def count_status(personnel):
    harvested = 0
    dormant = 0
    producing = 0
    for employee in personnel:
        status = employee.status.lower()
        match status:
            case "harvested":
                harvested += 1
            case "dormant":
                dormant += 1
            case "producing":
                producing +=1
    print(f'harvested: {harvested}\n dormant: {dormant} \n producing: {producing}') 
if __name__ == "__main__":
    main()