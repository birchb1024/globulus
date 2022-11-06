### NAME
`globulus`  -  Filter to re-order columns of a CSV file by entropy, 
ascending left to right. 

### SYNPOSIS

`globulus [filename]`

### DESCRIPTION

The program re-orders the columns of a CSV file given as the argument, or on
standard input if no file is specified. The aim is to produce a hierarchical arrangement
of the records. This is best understood by an example. Given this input

| Description                                        | Agent              | Price | Region  | Category                       |
| -------------------------------------------------- | ------------------ | ----- | ------- | ------------------------------ |
| Eldon Base for stackable storage shelf, platinum   | Muhammed MacIntyre | 35    | Nunavut | Storage & Organization         |
| 1.7 Cubic Foot Compact "Cube" Office Refrigerators | Barry French       | 68.02 | Nunavut | Appliances                     |
| Cardinal Slant-D® Ring Binder, Heavy Gauge Vinyl   | Barry French       | 2.99  | Nunavut | Binders and Binder Accessories |
| R380                                               | Clay Rozendal      | 3.99  | Nunavut | Telephones and Communication   |
| Holmes HEPA Air Purifier                           | Carlos Soltero     | 5.94  | Nunavut | Appliances                     |
| G.E. Longer-Life Indoor Recessed Floodlight Bulbs  | Carlos Soltero     | 4.95  | Nunavut | Office Furnishings             |
| Angle-D Binders with Locking Rings, Label Holders  | Carl Jackson       | 7.72  | Nunavut | Binders and Binder Accessories |
| SAFCO Mobile Desk Side File, Wire Frame            | Carl Jackson       | 6.22  | Nunavut | Storage & Organization         |
| SAFCO Commercial Wire Shelving, Black              | Monica Federle     | 35    | Nunavut | Storage & Organization         |
| Xerox 198                                          | Dorothy Badders    | 8.33  | Nunavut | Paper                          |

The program re-orders the columns like this:

| Region  | Category                       | Agent              | Price | Description                                        |
| ------- | ------------------------------ | ------------------ | ----- | -------------------------------------------------- |
| Nunavut | Appliances                     | Barry French       | 68.02 | 1.7 Cubic Foot Compact "Cube" Office Refrigerators |
| Nunavut | Appliances                     | Carlos Soltero     | 5.94  | Holmes HEPA Air Purifier                           |
| Nunavut | Binders and Binder Accessories | Barry French       | 2.99  | Cardinal Slant-D® Ring Binder, Heavy Gauge Vinyl   |
| Nunavut | Binders and Binder Accessories | Carl Jackson       | 7.72  | Angle-D Binders with Locking Rings, Label Holders  |
| Nunavut | Office Furnishings             | Carlos Soltero     | 4.95  | G.E. Longer-Life Indoor Recessed Floodlight Bulbs  |
| Nunavut | Paper                          | Dorothy Badders    | 8.33  | Xerox 198                                          |
| Nunavut | Storage & Organization         | Carl Jackson       | 6.22  | SAFCO Mobile Desk Side File, Wire Frame            |
| Nunavut | Storage & Organization         | Monica Federle     | 35    | SAFCO Commercial Wire Shelving, Black              |
| Nunavut | Storage & Organization         | Muhammed MacIntyre | 35    | Eldon Base for stackable storage shelf, platinum   |
| Nunavut | Telephones and Communication   | Clay Rozendal      | 3.99  | R380                                               |

It works in these steps:

1. counts the number distinct values in each column
2. reorders the columns with priority given to those with the least distinct values
3. outputs the table in the new order.

### EXAMPLES

This is a good companion for [frangipanni](https://github.com/birchb1024/frangipanni). Which can take the output from
`globulus` and generate the hierarchy, for example we get:

```
$ ./globulus <test/fixtures/SampleCSVFile_2kb_5col.csv | frangipanni -csv 
Region,Category,Agent,Price,Description
Nunavut
    Storage & Organization
        Muhammed MacIntyre,35,Eldon Base for stackable storage shelf, platinum
        Carl Jackson,6.22,SAFCO Mobile Desk Side File, Wire Frame
        Monica Federle,35,SAFCO Commercial Wire Shelving, Black
    Appliances
        Barry French,68.02,1.7 Cubic Foot Compact "Cube" Office Refrigerators
        Carlos Soltero,5.94,Holmes HEPA Air Purifier
    Binders and Binder Accessories
        Barry French,2.99,Cardinal Slant-D® Ring Binder, Heavy Gauge Vinyl
        Carl Jackson,7.72,Angle-D Binders with Locking Rings, Label Holders
    Telephones and Communication,Clay Rozendal,3.99,R380
    Office Furnishings,Carlos Soltero,4.95,G.E. Longer-Life Indoor Recessed Floodlight Bulbs
    Paper,Dorothy Badders,8.33,Xerox 198

```

### AUTHOR
Written by Peter Birch

### REPORTING BUGS
Report any bugs via GitHub issues: <https://github.com/birchb1024/globulus/issues>

### COPYRIGHT
Copyright © Peter William Birch: License MIT

### SEE ALSO
[frangipanni](https://github.com/birchb1024/frangipanni)

#### Eucalyptus Globulus

![](eucalyptus-globulus.jpeg)


