# yqp - YAMN quoted-printable

## [usage]

## yqp -i infile.txt -o outfile.txt

yqp is a small utility program which allows [YAMN](https://github.com/crooks/yamn) users, using
international character sets with an UTF-8 compatible editor,
to convert their messages to quoted-printable messages.

The Subject: line will be converted to UTF-8 base64 and folded
if the Subject: line is very long and contains UTF-8 characters.

In case the Subject: line contains only 7bit ASCII characters,
like a [hsub](https://github.com/crooks/hsubgen), it will be not converted.

yqp automatically adds the following Headers to a message:

Mime-Version: 1.0, Content-Type: text/plain; charset=UTF-8, Content-Transfer-Encoding: quoted-printable

I hope that you find yqp, in combination with YAMN, useful!
