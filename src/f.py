def main():
    return "hi"

if __name__ == "__main__":
    print(main())

"""
f
 - get data
 - update buffer
 - apply current filter
 - update output buffer if necessary

 on filter change:
  - kill subprocess if necessary
  - clear output buffer
  - start new filter process for current data set
  - update display
  - if more data since update goto 0, but only new data and update the display
"""