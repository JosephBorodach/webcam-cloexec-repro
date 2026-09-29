module cloexec-repro

go 1.22

// Point at the patched fork containing the OpenCloexec variant.
// Replace with your own fork/branch if reviewing an alternate patch.
replace github.com/blackjack/webcam => github.com/JosephBorodach/webcam open-with-cloexec

require github.com/blackjack/webcam v0.0.0-00010101000000-000000000000
