# daydayup

Usage

``` shell
git clone https://github.com/luozui/daydayup.git && cd daydayup
git clone -b main https://github.com/situ2001/gzhu_no_clock_in.git --depth 1

sudo apt-get install tesseract-ocr nodejs  # on ubuntu or debian
cd gzhu_no_clock_in
sudo pip3 install -r requirements.txt

# Runing at 8:00 every day.
crontab -e
# add this line:
# 0 8 * * * XXXXXPATH/run.sh

./server -a ":8080"
# nohup ./server -a ":8080" >server.log 2>&1 &
# pkill -9 ./server #stop
```

More:

```
Usage of ./server:
  -a, --apiaddr string   Admin api address (default ":8080")
  -l, --logdir string    logdir (default "./out.log")
  -u, --uidsdir string   uidsdir (default "./gzhu_no_clock_in/stu_id.txt")
pflag: help requested
```
