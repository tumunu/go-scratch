These programs can be executed via 'go run' or 'dot slash'

# carpentry
A collection of tools for a techno-carpenter

## wall-calulator
A calculator that determines the amount of timber required when framing a basic wall (excludes wall junctions and openings).
Based on NZS 3604 (New Zealand Building Code)

Unloaded wall calculation
```bash
./wall-calculator -h 2700 -l 8000
Timber stock length (mm): 4800
 
# studs in wall: 13
# timber stock for top/bottom plates: 4
# timber stock for studs: 13
 
wastage: 1.7
Total # timber stock required: 19
```

Loaded wall calculation
```bash
./wall-calculator -h 2700 -l 8000 -w true
Timber stock length (mm): 4800
 
# studs in wall: 20
# timber stock for top/bottom plates: 4
# timber stock for studs: 20
 
wastage: 2.4
Total # timber stock required: 26

```
