.PHONY: screenshots clean

cmd/img2ascii:
	go build -o cmd/img2ascii cmd/img2ascii.go

screenshots: cmd/img2ascii
	xfce4-terminal --working-directory=`pwd` --geometry 80x17 --hide-borders -H -e 'cmd/img2ascii --converter=ascii --width=80 images/sample.png' && sleep 1 && xfce4-screenshooter -w -s images/ascii.png
	xfce4-terminal --working-directory=`pwd` --geometry 80x17 --hide-borders -H -e 'cmd/img2ascii --converter=ansi --width=80 images/sample.png' && sleep 1 && xfce4-screenshooter -w -s images/ansi.png
	xfce4-terminal --working-directory=`pwd` --geometry 80x17 --hide-borders -H -e 'cmd/img2ascii --converter=ansi2562x --width=80 images/sample.png' && sleep 1 && xfce4-screenshooter -w -s images/ansiHalf.png
	xfce4-terminal --working-directory=`pwd` --geometry 80x17 --hide-borders -H -e 'cmd/img2ascii --converter=24bit --width=80 images/sample.png' && sleep 1 && xfce4-screenshooter -w -s images/24bit.png
	xfce4-terminal --working-directory=`pwd` --geometry 80x17 --hide-borders -H -e 'cmd/img2ascii --converter=24bit2x --width=80 images/sample.png' && sleep 1 && xfce4-screenshooter -w -s images/24bitHalf.png

clean:
	rm cmd/img2ascii