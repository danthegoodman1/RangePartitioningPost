# RangePartitioningPost

The worst memory DB using range-partitioning.

See the post at www.aspiring.dev/range-partitioning/

Run with:

```
go run .
```

Output:

```
Running range splitting!
Current DB:
Range:
	Low: ''
	High: 'Inf'
	Size: 500


### Range too large, splitting! ###

Current DB:
Range:
	Low: '00300'
	High: 'Inf'
	Size: 400
Range:
	Low: ''
	High: '00300'
	Size: 300


### Range too large, splitting! ###


### Range too large, splitting! ###


### Range too large, splitting! ###


### Range too large, splitting! ###

Current DB:
Range:
	Low: '01500'
	High: 'Inf'
	Size: 499
Range:
	Low: ''
	High: '00300'
	Size: 300
Range:
	Low: '00300'
	High: '00600'
	Size: 300
Range:
	Low: '00600'
	High: '00900'
	Size: 300
Range:
	Low: '00900'
	High: '01200'
	Size: 300
Range:
	Low: '01200'
	High: '01500'
	Size: 300

Getting key 01000 from range with low 00900
01000

### Range too large, splitting! ###


### Range too large, splitting! ###


### Range too large, splitting! ###


### Range too large, splitting! ###


### Range too large, splitting! ###


### Range too large, splitting! ###


### Range too large, splitting! ###


### Range too large, splitting! ###

Range:
	Low: '5497'
	High: 'Inf'
	Size: 5001
Range:
	Low: '41497'
	High: '5497'
	Size: 9999
Range:
	Low: ''
	High: '14497'
	Size: 5000
Range:
	Low: '14497'
	High: '18998'
	Size: 5000
Range:
	Low: '18998'
	High: '23497'
	Size: 5000
Range:
	Low: '23497'
	High: '27998'
	Size: 5000
Range:
	Low: '27998'
	High: '32497'
	Size: 5000
Range:
	Low: '32497'
	High: '36998'
	Size: 5000
Range:
	Low: '36998'
	High: '41497'
	Size: 5000

Getting key 42000 from range with low 41497
e
Getting key 20000 from range with low 18998
e
```