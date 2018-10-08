use v6;
use NativeCall;
use Test;

sub alloc_message(Str)
    returns Pointer[Str]
    is native('clang/lib/libmessage.so') { * }
sub free_message(Pointer[Str])
    is native('clang/lib/libmessage.so') { * }

subtest {
    for 1 .. 10 {
        my $input = "hogehoge{$_}";
        my $ptr = alloc_message($input);
        my $actual = $ptr.deref;
        free_message($ptr);

        is $actual, "Hello, {$input}", "Hello, {$input}";
    }
}, 'alloc_message';

done-testing;
