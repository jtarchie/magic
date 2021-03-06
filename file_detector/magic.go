
//line magic.rl:1
package file_detector

type Type int

const (
  //archive
  Epub Type = iota
  Zip
  Tar
  Rar
  Gzip
  Bz2
  SevenZip
  Pdf
  Exe
  Swf
  Rtf
  Nes
  Crs
  Cab
  Eot
  Ps
  Xz
  Deb
  Ar
  Z
  Lz
  Rpm
  Elf
  Dcm

  //application
  Wasm

  //audio
  Midi
  Mp3
  M4a
  Ogg
  Flac
  Wav
  Amr
  Aac

  //image
  Jpeg
  Jpeg2000
  Png
  Gif
  Webp
  CR2
  Tiff
  Bmp
  Jxr
  Psd
  Ico

  //unknown
  Unknown
)


//line magic.go:66
const magic_start int = 1
const magic_first_final int = 391
const magic_error int = 0

const magic_en_main int = 1


//line magic.rl:122


func Detect(data []byte) Type {
  cs, p, pe, eof := 0, 0, len(data), len(data)

//line magic.go:80
	{
	cs = magic_start
	}

//line magic.go:85
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 143:
		goto st_case_143
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	case 162:
		goto st_case_162
	case 163:
		goto st_case_163
	case 164:
		goto st_case_164
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 168:
		goto st_case_168
	case 169:
		goto st_case_169
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 175:
		goto st_case_175
	case 176:
		goto st_case_176
	case 177:
		goto st_case_177
	case 178:
		goto st_case_178
	case 179:
		goto st_case_179
	case 180:
		goto st_case_180
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
	case 183:
		goto st_case_183
	case 184:
		goto st_case_184
	case 185:
		goto st_case_185
	case 186:
		goto st_case_186
	case 187:
		goto st_case_187
	case 188:
		goto st_case_188
	case 189:
		goto st_case_189
	case 190:
		goto st_case_190
	case 191:
		goto st_case_191
	case 192:
		goto st_case_192
	case 193:
		goto st_case_193
	case 194:
		goto st_case_194
	case 195:
		goto st_case_195
	case 196:
		goto st_case_196
	case 197:
		goto st_case_197
	case 198:
		goto st_case_198
	case 199:
		goto st_case_199
	case 200:
		goto st_case_200
	case 201:
		goto st_case_201
	case 202:
		goto st_case_202
	case 203:
		goto st_case_203
	case 204:
		goto st_case_204
	case 205:
		goto st_case_205
	case 206:
		goto st_case_206
	case 207:
		goto st_case_207
	case 208:
		goto st_case_208
	case 209:
		goto st_case_209
	case 210:
		goto st_case_210
	case 211:
		goto st_case_211
	case 212:
		goto st_case_212
	case 213:
		goto st_case_213
	case 214:
		goto st_case_214
	case 215:
		goto st_case_215
	case 216:
		goto st_case_216
	case 217:
		goto st_case_217
	case 218:
		goto st_case_218
	case 219:
		goto st_case_219
	case 220:
		goto st_case_220
	case 221:
		goto st_case_221
	case 222:
		goto st_case_222
	case 223:
		goto st_case_223
	case 224:
		goto st_case_224
	case 225:
		goto st_case_225
	case 226:
		goto st_case_226
	case 227:
		goto st_case_227
	case 228:
		goto st_case_228
	case 229:
		goto st_case_229
	case 230:
		goto st_case_230
	case 231:
		goto st_case_231
	case 232:
		goto st_case_232
	case 233:
		goto st_case_233
	case 234:
		goto st_case_234
	case 235:
		goto st_case_235
	case 236:
		goto st_case_236
	case 237:
		goto st_case_237
	case 238:
		goto st_case_238
	case 239:
		goto st_case_239
	case 240:
		goto st_case_240
	case 241:
		goto st_case_241
	case 242:
		goto st_case_242
	case 243:
		goto st_case_243
	case 244:
		goto st_case_244
	case 245:
		goto st_case_245
	case 246:
		goto st_case_246
	case 247:
		goto st_case_247
	case 248:
		goto st_case_248
	case 249:
		goto st_case_249
	case 250:
		goto st_case_250
	case 251:
		goto st_case_251
	case 252:
		goto st_case_252
	case 253:
		goto st_case_253
	case 254:
		goto st_case_254
	case 255:
		goto st_case_255
	case 256:
		goto st_case_256
	case 257:
		goto st_case_257
	case 258:
		goto st_case_258
	case 0:
		goto st_case_0
	case 259:
		goto st_case_259
	case 260:
		goto st_case_260
	case 261:
		goto st_case_261
	case 262:
		goto st_case_262
	case 391:
		goto st_case_391
	case 392:
		goto st_case_392
	case 263:
		goto st_case_263
	case 264:
		goto st_case_264
	case 265:
		goto st_case_265
	case 266:
		goto st_case_266
	case 267:
		goto st_case_267
	case 268:
		goto st_case_268
	case 269:
		goto st_case_269
	case 270:
		goto st_case_270
	case 271:
		goto st_case_271
	case 393:
		goto st_case_393
	case 394:
		goto st_case_394
	case 395:
		goto st_case_395
	case 396:
		goto st_case_396
	case 397:
		goto st_case_397
	case 398:
		goto st_case_398
	case 399:
		goto st_case_399
	case 400:
		goto st_case_400
	case 401:
		goto st_case_401
	case 402:
		goto st_case_402
	case 403:
		goto st_case_403
	case 404:
		goto st_case_404
	case 405:
		goto st_case_405
	case 406:
		goto st_case_406
	case 407:
		goto st_case_407
	case 408:
		goto st_case_408
	case 409:
		goto st_case_409
	case 410:
		goto st_case_410
	case 411:
		goto st_case_411
	case 412:
		goto st_case_412
	case 413:
		goto st_case_413
	case 414:
		goto st_case_414
	case 415:
		goto st_case_415
	case 416:
		goto st_case_416
	case 417:
		goto st_case_417
	case 418:
		goto st_case_418
	case 419:
		goto st_case_419
	case 420:
		goto st_case_420
	case 421:
		goto st_case_421
	case 422:
		goto st_case_422
	case 423:
		goto st_case_423
	case 424:
		goto st_case_424
	case 425:
		goto st_case_425
	case 426:
		goto st_case_426
	case 427:
		goto st_case_427
	case 428:
		goto st_case_428
	case 429:
		goto st_case_429
	case 430:
		goto st_case_430
	case 431:
		goto st_case_431
	case 432:
		goto st_case_432
	case 433:
		goto st_case_433
	case 434:
		goto st_case_434
	case 435:
		goto st_case_435
	case 436:
		goto st_case_436
	case 437:
		goto st_case_437
	case 438:
		goto st_case_438
	case 439:
		goto st_case_439
	case 440:
		goto st_case_440
	case 441:
		goto st_case_441
	case 442:
		goto st_case_442
	case 443:
		goto st_case_443
	case 444:
		goto st_case_444
	case 445:
		goto st_case_445
	case 446:
		goto st_case_446
	case 447:
		goto st_case_447
	case 448:
		goto st_case_448
	case 449:
		goto st_case_449
	case 450:
		goto st_case_450
	case 451:
		goto st_case_451
	case 452:
		goto st_case_452
	case 453:
		goto st_case_453
	case 454:
		goto st_case_454
	case 455:
		goto st_case_455
	case 456:
		goto st_case_456
	case 457:
		goto st_case_457
	case 458:
		goto st_case_458
	case 459:
		goto st_case_459
	case 460:
		goto st_case_460
	case 461:
		goto st_case_461
	case 462:
		goto st_case_462
	case 463:
		goto st_case_463
	case 464:
		goto st_case_464
	case 465:
		goto st_case_465
	case 466:
		goto st_case_466
	case 467:
		goto st_case_467
	case 468:
		goto st_case_468
	case 469:
		goto st_case_469
	case 470:
		goto st_case_470
	case 471:
		goto st_case_471
	case 472:
		goto st_case_472
	case 473:
		goto st_case_473
	case 474:
		goto st_case_474
	case 475:
		goto st_case_475
	case 476:
		goto st_case_476
	case 477:
		goto st_case_477
	case 478:
		goto st_case_478
	case 479:
		goto st_case_479
	case 480:
		goto st_case_480
	case 481:
		goto st_case_481
	case 482:
		goto st_case_482
	case 483:
		goto st_case_483
	case 484:
		goto st_case_484
	case 485:
		goto st_case_485
	case 486:
		goto st_case_486
	case 487:
		goto st_case_487
	case 488:
		goto st_case_488
	case 489:
		goto st_case_489
	case 490:
		goto st_case_490
	case 491:
		goto st_case_491
	case 492:
		goto st_case_492
	case 493:
		goto st_case_493
	case 494:
		goto st_case_494
	case 495:
		goto st_case_495
	case 496:
		goto st_case_496
	case 497:
		goto st_case_497
	case 498:
		goto st_case_498
	case 499:
		goto st_case_499
	case 500:
		goto st_case_500
	case 501:
		goto st_case_501
	case 502:
		goto st_case_502
	case 503:
		goto st_case_503
	case 504:
		goto st_case_504
	case 505:
		goto st_case_505
	case 506:
		goto st_case_506
	case 507:
		goto st_case_507
	case 508:
		goto st_case_508
	case 509:
		goto st_case_509
	case 510:
		goto st_case_510
	case 511:
		goto st_case_511
	case 512:
		goto st_case_512
	case 513:
		goto st_case_513
	case 514:
		goto st_case_514
	case 515:
		goto st_case_515
	case 516:
		goto st_case_516
	case 517:
		goto st_case_517
	case 518:
		goto st_case_518
	case 519:
		goto st_case_519
	case 520:
		goto st_case_520
	case 521:
		goto st_case_521
	case 522:
		goto st_case_522
	case 523:
		goto st_case_523
	case 524:
		goto st_case_524
	case 525:
		goto st_case_525
	case 526:
		goto st_case_526
	case 527:
		goto st_case_527
	case 528:
		goto st_case_528
	case 529:
		goto st_case_529
	case 530:
		goto st_case_530
	case 531:
		goto st_case_531
	case 532:
		goto st_case_532
	case 533:
		goto st_case_533
	case 534:
		goto st_case_534
	case 535:
		goto st_case_535
	case 536:
		goto st_case_536
	case 537:
		goto st_case_537
	case 538:
		goto st_case_538
	case 539:
		goto st_case_539
	case 540:
		goto st_case_540
	case 541:
		goto st_case_541
	case 542:
		goto st_case_542
	case 543:
		goto st_case_543
	case 544:
		goto st_case_544
	case 545:
		goto st_case_545
	case 546:
		goto st_case_546
	case 547:
		goto st_case_547
	case 548:
		goto st_case_548
	case 549:
		goto st_case_549
	case 550:
		goto st_case_550
	case 551:
		goto st_case_551
	case 552:
		goto st_case_552
	case 553:
		goto st_case_553
	case 554:
		goto st_case_554
	case 555:
		goto st_case_555
	case 556:
		goto st_case_556
	case 557:
		goto st_case_557
	case 558:
		goto st_case_558
	case 559:
		goto st_case_559
	case 560:
		goto st_case_560
	case 561:
		goto st_case_561
	case 562:
		goto st_case_562
	case 563:
		goto st_case_563
	case 564:
		goto st_case_564
	case 565:
		goto st_case_565
	case 566:
		goto st_case_566
	case 567:
		goto st_case_567
	case 568:
		goto st_case_568
	case 569:
		goto st_case_569
	case 570:
		goto st_case_570
	case 571:
		goto st_case_571
	case 572:
		goto st_case_572
	case 573:
		goto st_case_573
	case 574:
		goto st_case_574
	case 575:
		goto st_case_575
	case 576:
		goto st_case_576
	case 577:
		goto st_case_577
	case 578:
		goto st_case_578
	case 579:
		goto st_case_579
	case 580:
		goto st_case_580
	case 581:
		goto st_case_581
	case 582:
		goto st_case_582
	case 583:
		goto st_case_583
	case 584:
		goto st_case_584
	case 585:
		goto st_case_585
	case 586:
		goto st_case_586
	case 587:
		goto st_case_587
	case 588:
		goto st_case_588
	case 589:
		goto st_case_589
	case 590:
		goto st_case_590
	case 591:
		goto st_case_591
	case 592:
		goto st_case_592
	case 593:
		goto st_case_593
	case 594:
		goto st_case_594
	case 595:
		goto st_case_595
	case 596:
		goto st_case_596
	case 597:
		goto st_case_597
	case 598:
		goto st_case_598
	case 599:
		goto st_case_599
	case 600:
		goto st_case_600
	case 601:
		goto st_case_601
	case 602:
		goto st_case_602
	case 603:
		goto st_case_603
	case 604:
		goto st_case_604
	case 605:
		goto st_case_605
	case 606:
		goto st_case_606
	case 607:
		goto st_case_607
	case 608:
		goto st_case_608
	case 609:
		goto st_case_609
	case 610:
		goto st_case_610
	case 611:
		goto st_case_611
	case 612:
		goto st_case_612
	case 613:
		goto st_case_613
	case 614:
		goto st_case_614
	case 615:
		goto st_case_615
	case 616:
		goto st_case_616
	case 617:
		goto st_case_617
	case 618:
		goto st_case_618
	case 619:
		goto st_case_619
	case 620:
		goto st_case_620
	case 621:
		goto st_case_621
	case 622:
		goto st_case_622
	case 623:
		goto st_case_623
	case 624:
		goto st_case_624
	case 625:
		goto st_case_625
	case 626:
		goto st_case_626
	case 627:
		goto st_case_627
	case 628:
		goto st_case_628
	case 629:
		goto st_case_629
	case 630:
		goto st_case_630
	case 631:
		goto st_case_631
	case 632:
		goto st_case_632
	case 633:
		goto st_case_633
	case 634:
		goto st_case_634
	case 635:
		goto st_case_635
	case 636:
		goto st_case_636
	case 637:
		goto st_case_637
	case 638:
		goto st_case_638
	case 639:
		goto st_case_639
	case 640:
		goto st_case_640
	case 641:
		goto st_case_641
	case 272:
		goto st_case_272
	case 642:
		goto st_case_642
	case 643:
		goto st_case_643
	case 644:
		goto st_case_644
	case 645:
		goto st_case_645
	case 646:
		goto st_case_646
	case 647:
		goto st_case_647
	case 648:
		goto st_case_648
	case 649:
		goto st_case_649
	case 650:
		goto st_case_650
	case 651:
		goto st_case_651
	case 273:
		goto st_case_273
	case 274:
		goto st_case_274
	case 275:
		goto st_case_275
	case 276:
		goto st_case_276
	case 277:
		goto st_case_277
	case 278:
		goto st_case_278
	case 279:
		goto st_case_279
	case 280:
		goto st_case_280
	case 652:
		goto st_case_652
	case 281:
		goto st_case_281
	case 282:
		goto st_case_282
	case 283:
		goto st_case_283
	case 653:
		goto st_case_653
	case 654:
		goto st_case_654
	case 655:
		goto st_case_655
	case 656:
		goto st_case_656
	case 284:
		goto st_case_284
	case 285:
		goto st_case_285
	case 286:
		goto st_case_286
	case 287:
		goto st_case_287
	case 288:
		goto st_case_288
	case 289:
		goto st_case_289
	case 657:
		goto st_case_657
	case 658:
		goto st_case_658
	case 659:
		goto st_case_659
	case 660:
		goto st_case_660
	case 661:
		goto st_case_661
	case 662:
		goto st_case_662
	case 663:
		goto st_case_663
	case 664:
		goto st_case_664
	case 665:
		goto st_case_665
	case 666:
		goto st_case_666
	case 667:
		goto st_case_667
	case 668:
		goto st_case_668
	case 669:
		goto st_case_669
	case 670:
		goto st_case_670
	case 671:
		goto st_case_671
	case 290:
		goto st_case_290
	case 291:
		goto st_case_291
	case 292:
		goto st_case_292
	case 293:
		goto st_case_293
	case 294:
		goto st_case_294
	case 672:
		goto st_case_672
	case 295:
		goto st_case_295
	case 673:
		goto st_case_673
	case 296:
		goto st_case_296
	case 297:
		goto st_case_297
	case 674:
		goto st_case_674
	case 298:
		goto st_case_298
	case 299:
		goto st_case_299
	case 300:
		goto st_case_300
	case 301:
		goto st_case_301
	case 302:
		goto st_case_302
	case 675:
		goto st_case_675
	case 303:
		goto st_case_303
	case 304:
		goto st_case_304
	case 305:
		goto st_case_305
	case 676:
		goto st_case_676
	case 306:
		goto st_case_306
	case 677:
		goto st_case_677
	case 307:
		goto st_case_307
	case 678:
		goto st_case_678
	case 308:
		goto st_case_308
	case 309:
		goto st_case_309
	case 679:
		goto st_case_679
	case 310:
		goto st_case_310
	case 311:
		goto st_case_311
	case 680:
		goto st_case_680
	case 312:
		goto st_case_312
	case 313:
		goto st_case_313
	case 314:
		goto st_case_314
	case 681:
		goto st_case_681
	case 315:
		goto st_case_315
	case 316:
		goto st_case_316
	case 317:
		goto st_case_317
	case 682:
		goto st_case_682
	case 318:
		goto st_case_318
	case 319:
		goto st_case_319
	case 320:
		goto st_case_320
	case 321:
		goto st_case_321
	case 683:
		goto st_case_683
	case 684:
		goto st_case_684
	case 685:
		goto st_case_685
	case 686:
		goto st_case_686
	case 322:
		goto st_case_322
	case 323:
		goto st_case_323
	case 687:
		goto st_case_687
	case 324:
		goto st_case_324
	case 325:
		goto st_case_325
	case 326:
		goto st_case_326
	case 327:
		goto st_case_327
	case 688:
		goto st_case_688
	case 328:
		goto st_case_328
	case 329:
		goto st_case_329
	case 330:
		goto st_case_330
	case 331:
		goto st_case_331
	case 689:
		goto st_case_689
	case 332:
		goto st_case_332
	case 333:
		goto st_case_333
	case 334:
		goto st_case_334
	case 690:
		goto st_case_690
	case 335:
		goto st_case_335
	case 336:
		goto st_case_336
	case 337:
		goto st_case_337
	case 338:
		goto st_case_338
	case 339:
		goto st_case_339
	case 340:
		goto st_case_340
	case 691:
		goto st_case_691
	case 692:
		goto st_case_692
	case 341:
		goto st_case_341
	case 342:
		goto st_case_342
	case 343:
		goto st_case_343
	case 693:
		goto st_case_693
	case 344:
		goto st_case_344
	case 345:
		goto st_case_345
	case 346:
		goto st_case_346
	case 694:
		goto st_case_694
	case 695:
		goto st_case_695
	case 696:
		goto st_case_696
	case 697:
		goto st_case_697
	case 698:
		goto st_case_698
	case 699:
		goto st_case_699
	case 700:
		goto st_case_700
	case 701:
		goto st_case_701
	case 702:
		goto st_case_702
	case 703:
		goto st_case_703
	case 704:
		goto st_case_704
	case 705:
		goto st_case_705
	case 706:
		goto st_case_706
	case 707:
		goto st_case_707
	case 708:
		goto st_case_708
	case 709:
		goto st_case_709
	case 710:
		goto st_case_710
	case 711:
		goto st_case_711
	case 712:
		goto st_case_712
	case 713:
		goto st_case_713
	case 714:
		goto st_case_714
	case 715:
		goto st_case_715
	case 716:
		goto st_case_716
	case 717:
		goto st_case_717
	case 718:
		goto st_case_718
	case 719:
		goto st_case_719
	case 720:
		goto st_case_720
	case 721:
		goto st_case_721
	case 722:
		goto st_case_722
	case 723:
		goto st_case_723
	case 347:
		goto st_case_347
	case 348:
		goto st_case_348
	case 349:
		goto st_case_349
	case 350:
		goto st_case_350
	case 351:
		goto st_case_351
	case 352:
		goto st_case_352
	case 353:
		goto st_case_353
	case 354:
		goto st_case_354
	case 724:
		goto st_case_724
	case 355:
		goto st_case_355
	case 356:
		goto st_case_356
	case 357:
		goto st_case_357
	case 358:
		goto st_case_358
	case 359:
		goto st_case_359
	case 725:
		goto st_case_725
	case 360:
		goto st_case_360
	case 361:
		goto st_case_361
	case 362:
		goto st_case_362
	case 726:
		goto st_case_726
	case 363:
		goto st_case_363
	case 364:
		goto st_case_364
	case 365:
		goto st_case_365
	case 727:
		goto st_case_727
	case 366:
		goto st_case_366
	case 367:
		goto st_case_367
	case 368:
		goto st_case_368
	case 369:
		goto st_case_369
	case 370:
		goto st_case_370
	case 371:
		goto st_case_371
	case 372:
		goto st_case_372
	case 373:
		goto st_case_373
	case 374:
		goto st_case_374
	case 728:
		goto st_case_728
	case 375:
		goto st_case_375
	case 376:
		goto st_case_376
	case 377:
		goto st_case_377
	case 729:
		goto st_case_729
	case 378:
		goto st_case_378
	case 379:
		goto st_case_379
	case 380:
		goto st_case_380
	case 730:
		goto st_case_730
	case 381:
		goto st_case_381
	case 382:
		goto st_case_382
	case 383:
		goto st_case_383
	case 731:
		goto st_case_731
	case 384:
		goto st_case_384
	case 385:
		goto st_case_385
	case 386:
		goto st_case_386
	case 387:
		goto st_case_387
	case 388:
		goto st_case_388
	case 732:
		goto st_case_732
	case 389:
		goto st_case_389
	case 390:
		goto st_case_390
	case 733:
		goto st_case_733
	case 734:
		goto st_case_734
	case 735:
		goto st_case_735
	case 736:
		goto st_case_736
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 0:
			goto st2
		case 31:
			goto st282
		case 33:
			goto st284
		case 35:
			goto st290
		case 37:
			goto st295
		case 55:
			goto st298
		case 56:
			goto st303
		case 66:
			goto st306
		case 67:
			goto st308
		case 68:
			goto st312
		case 70:
			goto st315
		case 71:
			goto st316
		case 73:
			goto st318
		case 76:
			goto st324
		case 77:
			goto st332
		case 78:
			goto st341
		case 80:
			goto st344
		case 82:
			goto st348
		case 87:
			goto st360
		case 102:
			goto st363
		case 123:
			goto st371
		case 127:
			goto st375
		case 137:
			goto st378
		case 237:
			goto st381
		case 253:
			goto st384
		case 255:
			goto st389
		}
		goto st281
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 0:
			goto st3
		case 97:
			goto st275
		}
		goto st274
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 0:
			goto st4
		case 1:
			goto st272
		}
		goto st273
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 12 {
			goto st263
		}
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		goto st84
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		goto st100
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		goto st101
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		goto st109
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		goto st115
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		goto st117
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		goto st123
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		goto st125
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		goto st128
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		goto st129
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		goto st130
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		goto st131
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		goto st135
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		goto st137
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		goto st139
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		goto st140
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		goto st142
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		goto st143
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		goto st144
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		goto st145
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		goto st146
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		goto st147
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		goto st151
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		goto st152
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		goto st153
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		goto st154
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		goto st155
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		goto st156
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		goto st157
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		goto st158
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		goto st159
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		goto st160
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		goto st161
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		goto st162
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		goto st163
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		goto st164
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		goto st165
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		goto st166
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		goto st167
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		goto st168
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		goto st169
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		goto st170
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		goto st171
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		goto st172
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		goto st173
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		goto st174
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		goto st175
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		goto st176
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		goto st177
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		goto st178
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		goto st179
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		goto st180
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		goto st181
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		goto st182
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		goto st183
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		goto st184
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		goto st185
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		goto st186
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		goto st187
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		goto st188
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		goto st189
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		goto st190
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		goto st191
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		goto st192
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		goto st193
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		goto st194
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		goto st195
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		goto st196
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		goto st197
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		goto st198
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		goto st199
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		goto st200
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		goto st201
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		goto st202
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		goto st203
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		goto st204
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		goto st205
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		goto st206
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
		goto st207
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		goto st208
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		goto st209
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		goto st210
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		goto st211
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		goto st212
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		goto st213
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		goto st214
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		goto st215
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		goto st216
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		goto st217
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
		goto st218
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		goto st219
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		goto st220
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		goto st221
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		goto st222
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		goto st223
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		goto st224
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		goto st225
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		goto st226
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		goto st227
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		goto st228
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
		goto st230
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		goto st231
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		goto st232
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		goto st233
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
		goto st234
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		goto st235
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		goto st236
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		goto st237
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		goto st238
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		goto st239
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		goto st240
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		goto st241
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		goto st242
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		goto st243
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		goto st244
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		goto st245
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		goto st246
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		goto st247
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		goto st248
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		goto st249
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		goto st250
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		goto st251
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		goto st252
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		goto st253
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		goto st254
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		goto st255
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		goto st256
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		goto st257
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		goto st258
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		if data[p] == 117 {
			goto st259
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		if data[p] == 115 {
			goto st260
		}
		goto st0
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		if data[p] == 116 {
			goto st261
		}
		goto st0
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		if data[p] == 97 {
			goto st262
		}
		goto st0
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		if data[p] == 114 {
			goto st391
		}
		goto st0
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
		goto tr434
tr434:
//line magic.rl:68
return Tar;
	goto st392
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
//line magic.go:3238
		goto st392
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		if data[p] == 106 {
			goto st264
		}
		goto st6
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		if data[p] == 80 {
			goto st265
		}
		goto st7
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		if data[p] == 32 {
			goto st266
		}
		goto st8
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
		if data[p] == 32 {
			goto st267
		}
		goto st9
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		if data[p] == 13 {
			goto st268
		}
		goto st10
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		if data[p] == 10 {
			goto st269
		}
		goto st11
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		if data[p] == 135 {
			goto st270
		}
		goto st12
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		if data[p] == 10 {
			goto st271
		}
		goto st13
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		if data[p] == 0 {
			goto st393
		}
		goto st14
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
		goto tr436
tr436:
//line magic.rl:105
return Jpeg2000
	goto st394
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
//line magic.go:3336
		goto st395
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		goto st396
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
		goto st397
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		goto st398
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		goto st399
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		goto st400
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
		goto st401
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		goto st402
tr713:
//line magic.rl:83
return Deb
	goto st402
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
//line magic.go:3389
		goto st403
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		goto st404
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		goto st405
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		goto st406
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		goto st407
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		goto st408
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
		goto st409
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		goto st410
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		goto st411
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		goto st412
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		goto st413
tr766:
//line magic.rl:66
 return Epub;
	goto st413
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
//line magic.go:3460
		goto st414
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		goto st415
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		goto st416
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		goto st417
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		goto st418
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		goto st419
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		goto st420
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
		goto st421
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
		goto st422
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		goto st423
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
		goto st424
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		goto st425
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
		goto st426
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
		goto st427
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		goto st428
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
		goto st429
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
		goto st430
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
		goto st431
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		goto st432
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
		goto st433
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		goto st434
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
		goto st435
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		goto st436
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
		goto st437
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
		goto st438
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
		goto st439
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		goto st440
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
		goto st441
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
		goto st442
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
		goto st443
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
		goto st444
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		goto st445
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
		goto st446
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		goto st447
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
		goto st448
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
		goto st449
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		goto st450
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
		goto st451
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
		goto st452
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
		goto st453
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		goto st454
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		goto st455
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
		goto st456
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		goto st457
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		goto st458
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
		goto st459
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
		goto st460
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
		goto st461
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
		goto st462
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
		goto st463
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
		goto st464
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
		goto st465
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
		goto st466
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
		goto st467
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
		goto st468
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
		goto st469
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
		goto st470
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
		goto st471
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
		goto st472
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
		goto st473
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
		goto st474
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		goto st475
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
		goto st476
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
		goto st477
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		goto st478
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		goto st479
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
		goto st480
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
		goto st481
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
		goto st482
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
		goto st483
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
		goto st484
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
		goto st485
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
		goto st486
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		goto st487
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
		goto st488
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
		goto st489
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
		goto st490
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
		goto st491
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		goto st492
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
		goto st493
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
		goto st494
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		goto st495
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		goto st496
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
		goto st497
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		goto st498
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		goto st499
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		goto st500
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		goto st501
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		goto st502
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		goto st503
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		goto st504
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		goto st505
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		goto st506
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		goto st507
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		goto st508
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
		goto st509
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
		goto st510
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
		goto st511
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		goto st512
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
		goto st513
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
		goto st514
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
		goto st515
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
		goto st516
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
		goto st517
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
		goto st518
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
		goto st519
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		goto st520
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
		goto st521
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
		goto st522
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
		goto st523
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
		goto st524
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
		goto st525
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		goto st526
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
		goto st527
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
		goto st528
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
		goto st529
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
		goto st530
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
		goto st531
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
		goto st532
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
		goto st533
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
		goto st534
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
		goto st535
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		goto st536
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
		goto st537
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
		goto st538
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
		goto st539
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
		goto st540
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
		goto st541
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
		goto st542
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
		goto st543
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
		goto st544
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
		goto st545
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
		goto st546
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
		goto st547
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
		goto st548
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
		goto st549
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
		goto st550
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
		goto st551
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
		goto st552
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
		goto st553
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
		goto st554
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
		goto st555
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
		goto st556
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
		goto st557
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
		goto st558
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
		goto st559
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
		goto st560
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
		goto st561
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
		goto st562
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
		goto st563
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
		goto st564
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
		goto st565
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
		goto st566
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
		goto st567
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
		goto st568
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
		goto st569
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
		goto st570
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
		goto st571
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
		goto st572
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
		goto st573
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		goto st574
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
		goto st575
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
		goto st576
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
		goto st577
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
		goto st578
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
		goto st579
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
		goto st580
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
		goto st581
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
		goto st582
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
		goto st583
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
		goto st584
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
		goto st585
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		goto st586
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
	st_case_586:
		goto st587
	st587:
		if p++; p == pe {
			goto _test_eof587
		}
	st_case_587:
		goto st588
	st588:
		if p++; p == pe {
			goto _test_eof588
		}
	st_case_588:
		goto st589
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
		goto st590
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
		goto st591
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
		goto st592
	st592:
		if p++; p == pe {
			goto _test_eof592
		}
	st_case_592:
		goto st593
	st593:
		if p++; p == pe {
			goto _test_eof593
		}
	st_case_593:
		goto st594
	st594:
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
		goto st595
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
		goto st596
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		goto st597
	st597:
		if p++; p == pe {
			goto _test_eof597
		}
	st_case_597:
		goto st598
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
		goto st599
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		goto st600
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		goto st601
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
	st_case_601:
		goto st602
	st602:
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
		goto st603
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
		goto st604
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
		goto st605
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
		goto st606
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
		goto st607
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
		goto st608
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
		goto st609
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
		goto st610
	st610:
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
		goto st611
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		goto st612
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
		goto st613
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
		goto st614
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
		goto st615
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
		goto st616
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
		goto st617
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
		goto st618
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
		goto st619
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		goto st620
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
		goto st621
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		goto st622
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		goto st623
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
		goto st624
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
		goto st625
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
	st_case_625:
		goto st626
	st626:
		if p++; p == pe {
			goto _test_eof626
		}
	st_case_626:
		goto st627
	st627:
		if p++; p == pe {
			goto _test_eof627
		}
	st_case_627:
		goto st628
	st628:
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
		goto st629
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
		goto st630
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
		goto st631
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		goto st632
	st632:
		if p++; p == pe {
			goto _test_eof632
		}
	st_case_632:
		goto st633
	st633:
		if p++; p == pe {
			goto _test_eof633
		}
	st_case_633:
		goto st634
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
		goto st635
	st635:
		if p++; p == pe {
			goto _test_eof635
		}
	st_case_635:
		goto st636
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
		goto st637
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
	st_case_637:
		if data[p] == 117 {
			goto st638
		}
		goto st392
	st638:
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
		if data[p] == 115 {
			goto st639
		}
		goto st392
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
		if data[p] == 116 {
			goto st640
		}
		goto st392
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
	st_case_640:
		if data[p] == 97 {
			goto st641
		}
		goto st392
	st641:
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
		if data[p] == 114 {
			goto st391
		}
		goto st392
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		if data[p] == 0 {
			goto st642
		}
		goto st5
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
		goto tr684
tr684:
//line magic.rl:114
return Ico
	goto st643
tr716:
//line magic.rl:73
return Pdf
	goto st643
tr718:
//line magic.rl:113
return Psd
	goto st643
tr722:
//line magic.rl:78
return Crs
	goto st643
tr723:
//line magic.rl:89
return Dcm
	goto st643
tr725:
//line magic.rl:110
return Tiff
	goto st643
tr730:
//line magic.rl:79
return Cab
	goto st643
tr732:
//line magic.rl:86
return Lz
	goto st643
tr733:
//line magic.rl:97
return M4a
	goto st643
tr734:
//line magic.rl:95
return Midi
	goto st643
tr736:
//line magic.rl:77
return Nes
	goto st643
tr737:
//line magic.rl:67
return Zip;
	goto st643
tr769:
//line magic.rl:108
return Webp
	goto st643
tr770:
//line magic.rl:98
return Ogg
	goto st643
tr772:
//line magic.rl:88
return Elf
	goto st643
tr773:
//line magic.rl:106
return Png
	goto st643
tr774:
//line magic.rl:87
return Rpm
	goto st643
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
	st_case_643:
//line magic.go:4933
		goto st644
tr731:
//line magic.rl:80
return Eot
	goto st644
tr771:
//line magic.rl:76
return Rtf
	goto st644
	st644:
		if p++; p == pe {
			goto _test_eof644
		}
	st_case_644:
//line magic.go:4948
		goto st645
tr714:
//line magic.rl:100
return Amr
	goto st645
tr717:
//line magic.rl:72
return SevenZip;
	goto st645
tr728:
//line magic.rl:109
return CR2
	goto st645
tr775:
//line magic.rl:82
return Xz
	goto st645
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
//line magic.go:4971
		goto st646
tr698:
//line magic.rl:84
return Ar
	goto st646
tr768:
//line magic.rl:69
return Rar;
	goto st646
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
//line magic.go:4986
		goto st647
tr693:
//line magic.rl:92
return Wasm
	goto st647
tr767:
//line magic.rl:99
return Wav
	goto st647
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
//line magic.go:5001
		goto st648
	st648:
		if p++; p == pe {
			goto _test_eof648
		}
	st_case_648:
		goto st649
	st649:
		if p++; p == pe {
			goto _test_eof649
		}
	st_case_649:
		goto st650
	st650:
		if p++; p == pe {
			goto _test_eof650
		}
	st_case_650:
		goto st651
	st651:
		if p++; p == pe {
			goto _test_eof651
		}
	st_case_651:
		goto st394
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
		goto st5
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
		goto st273
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		if data[p] == 115 {
			goto st276
		}
		goto st273
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		if data[p] == 109 {
			goto st277
		}
		goto st5
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		if data[p] == 1 {
			goto st278
		}
		goto st6
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		if data[p] == 0 {
			goto st279
		}
		goto st7
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		if data[p] == 0 {
			goto st280
		}
		goto st8
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		if data[p] == 0 {
			goto st652
		}
		goto st9
	st652:
		if p++; p == pe {
			goto _test_eof652
		}
	st_case_652:
		goto tr693
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
		goto st274
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		switch data[p] {
		case 139:
			goto st283
		case 157:
			goto st655
		case 160:
			goto st655
		}
		goto st274
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		if data[p] == 8 {
			goto st653
		}
		goto st273
	st653:
		if p++; p == pe {
			goto _test_eof653
		}
	st_case_653:
		goto tr694
tr694:
//line magic.rl:70
return Gzip;
	goto st654
tr720:
//line magic.rl:71
return Bz2;
	goto st654
tr721:
//line magic.rl:75
return Swf
	goto st654
tr724:
//line magic.rl:107
return Gif
	goto st654
tr729:
//line magic.rl:112
return Jxr
	goto st654
tr776:
//line magic.rl:104
return Jpeg
	goto st654
	st654:
		if p++; p == pe {
			goto _test_eof654
		}
	st_case_654:
//line magic.go:5163
		goto st643
	st655:
		if p++; p == pe {
			goto _test_eof655
		}
	st_case_655:
		goto tr696
tr696:
//line magic.rl:85
return Z
	goto st656
tr715:
//line magic.rl:81
return Ps
	goto st656
tr719:
//line magic.rl:111
return Bmp
	goto st656
tr735:
//line magic.rl:74
return Exe
	goto st656
tr777:
//line magic.rl:101
return Aac
	goto st656
tr778:
//line magic.rl:96
return Mp3
	goto st656
	st656:
		if p++; p == pe {
			goto _test_eof656
		}
	st_case_656:
//line magic.go:5200
		goto st654
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		if data[p] == 60 {
			goto st285
		}
		goto st274
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		if data[p] == 97 {
			goto st286
		}
		goto st273
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		if data[p] == 114 {
			goto st287
		}
		goto st5
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		if data[p] == 99 {
			goto st288
		}
		goto st6
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
		if data[p] == 104 {
			goto st289
		}
		goto st7
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		if data[p] == 62 {
			goto st657
		}
		goto st8
	st657:
		if p++; p == pe {
			goto _test_eof657
		}
	st_case_657:
		if data[p] == 10 {
			goto tr699
		}
		goto tr698
tr699:
//line magic.rl:84
return Ar
	goto st658
	st658:
		if p++; p == pe {
			goto _test_eof658
		}
	st_case_658:
//line magic.go:5274
		if data[p] == 100 {
			goto st659
		}
		goto st647
	st659:
		if p++; p == pe {
			goto _test_eof659
		}
	st_case_659:
		if data[p] == 101 {
			goto st660
		}
		goto st648
	st660:
		if p++; p == pe {
			goto _test_eof660
		}
	st_case_660:
		if data[p] == 98 {
			goto st661
		}
		goto st649
	st661:
		if p++; p == pe {
			goto _test_eof661
		}
	st_case_661:
		if data[p] == 105 {
			goto st662
		}
		goto st650
	st662:
		if p++; p == pe {
			goto _test_eof662
		}
	st_case_662:
		if data[p] == 97 {
			goto st663
		}
		goto st651
	st663:
		if p++; p == pe {
			goto _test_eof663
		}
	st_case_663:
		if data[p] == 110 {
			goto st664
		}
		goto st394
	st664:
		if p++; p == pe {
			goto _test_eof664
		}
	st_case_664:
		if data[p] == 45 {
			goto st665
		}
		goto st395
	st665:
		if p++; p == pe {
			goto _test_eof665
		}
	st_case_665:
		if data[p] == 98 {
			goto st666
		}
		goto st396
	st666:
		if p++; p == pe {
			goto _test_eof666
		}
	st_case_666:
		if data[p] == 105 {
			goto st667
		}
		goto st397
	st667:
		if p++; p == pe {
			goto _test_eof667
		}
	st_case_667:
		if data[p] == 110 {
			goto st668
		}
		goto st398
	st668:
		if p++; p == pe {
			goto _test_eof668
		}
	st_case_668:
		if data[p] == 97 {
			goto st669
		}
		goto st399
	st669:
		if p++; p == pe {
			goto _test_eof669
		}
	st_case_669:
		if data[p] == 114 {
			goto st670
		}
		goto st400
	st670:
		if p++; p == pe {
			goto _test_eof670
		}
	st_case_670:
		if data[p] == 121 {
			goto st671
		}
		goto st401
	st671:
		if p++; p == pe {
			goto _test_eof671
		}
	st_case_671:
		goto tr713
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
		if data[p] == 33 {
			goto st291
		}
		goto st274
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		if data[p] == 65 {
			goto st292
		}
		goto st273
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
		if data[p] == 77 {
			goto st293
		}
		goto st5
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
		if data[p] == 82 {
			goto st294
		}
		goto st6
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		if data[p] == 10 {
			goto st672
		}
		goto st7
	st672:
		if p++; p == pe {
			goto _test_eof672
		}
	st_case_672:
		goto tr714
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		switch data[p] {
		case 33:
			goto st673
		case 80:
			goto st296
		}
		goto st274
	st673:
		if p++; p == pe {
			goto _test_eof673
		}
	st_case_673:
		goto tr715
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
		if data[p] == 68 {
			goto st297
		}
		goto st273
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
		if data[p] == 70 {
			goto st674
		}
		goto st5
	st674:
		if p++; p == pe {
			goto _test_eof674
		}
	st_case_674:
		goto tr716
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		if data[p] == 122 {
			goto st299
		}
		goto st274
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		if data[p] == 188 {
			goto st300
		}
		goto st273
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		if data[p] == 175 {
			goto st301
		}
		goto st5
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
		if data[p] == 39 {
			goto st302
		}
		goto st6
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		if data[p] == 28 {
			goto st675
		}
		goto st7
	st675:
		if p++; p == pe {
			goto _test_eof675
		}
	st_case_675:
		goto tr717
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		if data[p] == 66 {
			goto st304
		}
		goto st274
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		if data[p] == 80 {
			goto st305
		}
		goto st273
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		if data[p] == 83 {
			goto st676
		}
		goto st5
	st676:
		if p++; p == pe {
			goto _test_eof676
		}
	st_case_676:
		goto tr718
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		switch data[p] {
		case 77:
			goto st677
		case 90:
			goto st307
		}
		goto st274
	st677:
		if p++; p == pe {
			goto _test_eof677
		}
	st_case_677:
		goto tr719
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
		if data[p] == 104 {
			goto st678
		}
		goto st273
	st678:
		if p++; p == pe {
			goto _test_eof678
		}
	st_case_678:
		goto tr720
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
		switch data[p] {
		case 87:
			goto st309
		case 114:
			goto st310
		}
		goto st274
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		if data[p] == 83 {
			goto st679
		}
		goto st273
	st679:
		if p++; p == pe {
			goto _test_eof679
		}
	st_case_679:
		goto tr721
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		if data[p] == 50 {
			goto st311
		}
		goto st273
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
		if data[p] == 52 {
			goto st680
		}
		goto st5
	st680:
		if p++; p == pe {
			goto _test_eof680
		}
	st_case_680:
		goto tr722
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		if data[p] == 73 {
			goto st313
		}
		goto st274
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		if data[p] == 67 {
			goto st314
		}
		goto st273
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		if data[p] == 77 {
			goto st681
		}
		goto st5
	st681:
		if p++; p == pe {
			goto _test_eof681
		}
	st_case_681:
		goto tr723
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
		if data[p] == 87 {
			goto st309
		}
		goto st274
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
		if data[p] == 73 {
			goto st317
		}
		goto st274
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
		if data[p] == 70 {
			goto st682
		}
		goto st273
	st682:
		if p++; p == pe {
			goto _test_eof682
		}
	st_case_682:
		goto tr724
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		switch data[p] {
		case 68:
			goto st319
		case 73:
			goto st320
		case 83:
			goto st322
		}
		goto st274
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		if data[p] == 51 {
			goto st656
		}
		goto st273
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		switch data[p] {
		case 42:
			goto st321
		case 188:
			goto st686
		}
		goto st273
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
		if data[p] == 0 {
			goto st683
		}
		goto st5
	st683:
		if p++; p == pe {
			goto _test_eof683
		}
	st_case_683:
		if data[p] == 67 {
			goto tr726
		}
		goto tr725
tr726:
//line magic.rl:110
return Tiff
	goto st684
	st684:
		if p++; p == pe {
			goto _test_eof684
		}
	st_case_684:
//line magic.go:5782
		if data[p] == 82 {
			goto st685
		}
		goto st644
	st685:
		if p++; p == pe {
			goto _test_eof685
		}
	st_case_685:
		goto tr728
	st686:
		if p++; p == pe {
			goto _test_eof686
		}
	st_case_686:
		goto tr729
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		if data[p] == 99 {
			goto st323
		}
		goto st273
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		if data[p] == 40 {
			goto st687
		}
		goto st5
	st687:
		if p++; p == pe {
			goto _test_eof687
		}
	st_case_687:
		goto tr730
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		switch data[p] {
		case 80:
			goto st325
		case 90:
			goto st330
		}
		goto st274
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
		switch data[p] {
		case 1:
			goto st326
		case 2:
			goto st328
		}
		goto st273
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		if data[p] == 0 {
			goto st327
		}
		goto st5
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
		if data[p] == 0 {
			goto st688
		}
		goto st6
	st688:
		if p++; p == pe {
			goto _test_eof688
		}
	st_case_688:
		goto tr731
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		if data[p] == 0 {
			goto st329
		}
		goto st5
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
		if 1 <= data[p] && data[p] <= 2 {
			goto st688
		}
		goto st6
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		if data[p] == 73 {
			goto st331
		}
		goto st273
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		if data[p] == 80 {
			goto st689
		}
		goto st5
	st689:
		if p++; p == pe {
			goto _test_eof689
		}
	st_case_689:
		goto tr732
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		switch data[p] {
		case 52:
			goto st333
		case 77:
			goto st335
		case 83:
			goto st337
		case 84:
			goto st339
		case 90:
			goto st692
		}
		goto st274
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		if data[p] == 65 {
			goto st334
		}
		goto st273
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		if data[p] == 32 {
			goto st690
		}
		goto st5
	st690:
		if p++; p == pe {
			goto _test_eof690
		}
	st_case_690:
		goto tr733
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		if data[p] == 0 {
			goto st336
		}
		goto st273
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		if data[p] == 42 {
			goto st683
		}
		goto st5
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		if data[p] == 67 {
			goto st338
		}
		goto st273
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
		if data[p] == 70 {
			goto st687
		}
		goto st5
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		if data[p] == 104 {
			goto st340
		}
		goto st273
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		if data[p] == 100 {
			goto st691
		}
		goto st5
	st691:
		if p++; p == pe {
			goto _test_eof691
		}
	st_case_691:
		goto tr734
	st692:
		if p++; p == pe {
			goto _test_eof692
		}
	st_case_692:
		goto tr735
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		if data[p] == 69 {
			goto st342
		}
		goto st274
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
		if data[p] == 83 {
			goto st343
		}
		goto st273
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
		if data[p] == 26 {
			goto st693
		}
		goto st5
	st693:
		if p++; p == pe {
			goto _test_eof693
		}
	st_case_693:
		goto tr736
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
		if data[p] == 75 {
			goto st345
		}
		goto st274
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
		switch data[p] {
		case 3:
			goto st346
		case 5:
			goto st347
		case 7:
			goto st347
		}
		goto st273
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
		switch data[p] {
		case 4:
			goto st694
		case 6:
			goto st723
		case 8:
			goto st723
		}
		goto st5
	st694:
		if p++; p == pe {
			goto _test_eof694
		}
	st_case_694:
		if data[p] == 109 {
			goto tr738
		}
		goto tr737
tr738:
//line magic.rl:67
return Zip;
	goto st695
	st695:
		if p++; p == pe {
			goto _test_eof695
		}
	st_case_695:
//line magic.go:6109
		if data[p] == 105 {
			goto st696
		}
		goto st644
	st696:
		if p++; p == pe {
			goto _test_eof696
		}
	st_case_696:
		if data[p] == 109 {
			goto st697
		}
		goto st645
	st697:
		if p++; p == pe {
			goto _test_eof697
		}
	st_case_697:
		if data[p] == 101 {
			goto st698
		}
		goto st646
	st698:
		if p++; p == pe {
			goto _test_eof698
		}
	st_case_698:
		if data[p] == 116 {
			goto st699
		}
		goto st647
	st699:
		if p++; p == pe {
			goto _test_eof699
		}
	st_case_699:
		if data[p] == 121 {
			goto st700
		}
		goto st648
	st700:
		if p++; p == pe {
			goto _test_eof700
		}
	st_case_700:
		if data[p] == 112 {
			goto st701
		}
		goto st649
	st701:
		if p++; p == pe {
			goto _test_eof701
		}
	st_case_701:
		if data[p] == 101 {
			goto st702
		}
		goto st650
	st702:
		if p++; p == pe {
			goto _test_eof702
		}
	st_case_702:
		if data[p] == 97 {
			goto st703
		}
		goto st651
	st703:
		if p++; p == pe {
			goto _test_eof703
		}
	st_case_703:
		if data[p] == 112 {
			goto st704
		}
		goto st394
	st704:
		if p++; p == pe {
			goto _test_eof704
		}
	st_case_704:
		if data[p] == 112 {
			goto st705
		}
		goto st395
	st705:
		if p++; p == pe {
			goto _test_eof705
		}
	st_case_705:
		if data[p] == 108 {
			goto st706
		}
		goto st396
	st706:
		if p++; p == pe {
			goto _test_eof706
		}
	st_case_706:
		if data[p] == 105 {
			goto st707
		}
		goto st397
	st707:
		if p++; p == pe {
			goto _test_eof707
		}
	st_case_707:
		if data[p] == 99 {
			goto st708
		}
		goto st398
	st708:
		if p++; p == pe {
			goto _test_eof708
		}
	st_case_708:
		if data[p] == 97 {
			goto st709
		}
		goto st399
	st709:
		if p++; p == pe {
			goto _test_eof709
		}
	st_case_709:
		if data[p] == 116 {
			goto st710
		}
		goto st400
	st710:
		if p++; p == pe {
			goto _test_eof710
		}
	st_case_710:
		if data[p] == 105 {
			goto st711
		}
		goto st401
	st711:
		if p++; p == pe {
			goto _test_eof711
		}
	st_case_711:
		if data[p] == 111 {
			goto st712
		}
		goto st402
	st712:
		if p++; p == pe {
			goto _test_eof712
		}
	st_case_712:
		if data[p] == 110 {
			goto st713
		}
		goto st403
	st713:
		if p++; p == pe {
			goto _test_eof713
		}
	st_case_713:
		if data[p] == 47 {
			goto st714
		}
		goto st404
	st714:
		if p++; p == pe {
			goto _test_eof714
		}
	st_case_714:
		if data[p] == 101 {
			goto st715
		}
		goto st405
	st715:
		if p++; p == pe {
			goto _test_eof715
		}
	st_case_715:
		if data[p] == 112 {
			goto st716
		}
		goto st406
	st716:
		if p++; p == pe {
			goto _test_eof716
		}
	st_case_716:
		if data[p] == 117 {
			goto st717
		}
		goto st407
	st717:
		if p++; p == pe {
			goto _test_eof717
		}
	st_case_717:
		if data[p] == 98 {
			goto st718
		}
		goto st408
	st718:
		if p++; p == pe {
			goto _test_eof718
		}
	st_case_718:
		if data[p] == 43 {
			goto st719
		}
		goto st409
	st719:
		if p++; p == pe {
			goto _test_eof719
		}
	st_case_719:
		if data[p] == 122 {
			goto st720
		}
		goto st410
	st720:
		if p++; p == pe {
			goto _test_eof720
		}
	st_case_720:
		if data[p] == 105 {
			goto st721
		}
		goto st411
	st721:
		if p++; p == pe {
			goto _test_eof721
		}
	st_case_721:
		if data[p] == 112 {
			goto st722
		}
		goto st412
	st722:
		if p++; p == pe {
			goto _test_eof722
		}
	st_case_722:
		goto tr766
	st723:
		if p++; p == pe {
			goto _test_eof723
		}
	st_case_723:
		goto tr737
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
		switch data[p] {
		case 4:
			goto st723
		case 6:
			goto st723
		case 8:
			goto st723
		}
		goto st5
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		switch data[p] {
		case 73:
			goto st349
		case 97:
			goto st355
		}
		goto st274
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		if data[p] == 70 {
			goto st350
		}
		goto st273
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		if data[p] == 70 {
			goto st351
		}
		goto st5
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		if data[p] == 87 {
			goto st352
		}
		goto st6
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		if data[p] == 65 {
			goto st353
		}
		goto st7
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		if data[p] == 86 {
			goto st354
		}
		goto st8
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		if data[p] == 69 {
			goto st724
		}
		goto st9
	st724:
		if p++; p == pe {
			goto _test_eof724
		}
	st_case_724:
		goto tr767
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		if data[p] == 114 {
			goto st356
		}
		goto st273
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		if data[p] == 33 {
			goto st357
		}
		goto st5
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		if data[p] == 26 {
			goto st358
		}
		goto st6
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
		if data[p] == 7 {
			goto st359
		}
		goto st7
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
		if data[p] <= 1 {
			goto st725
		}
		goto st8
	st725:
		if p++; p == pe {
			goto _test_eof725
		}
	st_case_725:
		goto tr768
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
		if data[p] == 69 {
			goto st361
		}
		goto st274
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		if data[p] == 66 {
			goto st362
		}
		goto st273
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		if data[p] == 80 {
			goto st726
		}
		goto st5
	st726:
		if p++; p == pe {
			goto _test_eof726
		}
	st_case_726:
		goto tr769
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		switch data[p] {
		case 76:
			goto st364
		case 116:
			goto st366
		}
		goto st274
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		if data[p] == 97 {
			goto st365
		}
		goto st273
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		if data[p] == 67 {
			goto st727
		}
		goto st5
	st727:
		if p++; p == pe {
			goto _test_eof727
		}
	st_case_727:
		goto tr770
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		if data[p] == 121 {
			goto st367
		}
		goto st273
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		if data[p] == 112 {
			goto st368
		}
		goto st5
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		if data[p] == 77 {
			goto st369
		}
		goto st6
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
		if data[p] == 52 {
			goto st370
		}
		goto st7
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		if data[p] == 65 {
			goto st645
		}
		goto st8
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		if data[p] == 92 {
			goto st372
		}
		goto st274
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		if data[p] == 114 {
			goto st373
		}
		goto st273
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
		if data[p] == 116 {
			goto st374
		}
		goto st5
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
		if data[p] == 102 {
			goto st728
		}
		goto st6
	st728:
		if p++; p == pe {
			goto _test_eof728
		}
	st_case_728:
		goto tr771
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		if data[p] == 69 {
			goto st376
		}
		goto st274
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		if data[p] == 76 {
			goto st377
		}
		goto st273
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		if data[p] == 70 {
			goto st729
		}
		goto st5
	st729:
		if p++; p == pe {
			goto _test_eof729
		}
	st_case_729:
		goto tr772
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		if data[p] == 80 {
			goto st379
		}
		goto st274
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		if data[p] == 78 {
			goto st380
		}
		goto st273
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		if data[p] == 71 {
			goto st730
		}
		goto st5
	st730:
		if p++; p == pe {
			goto _test_eof730
		}
	st_case_730:
		goto tr773
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		if data[p] == 171 {
			goto st382
		}
		goto st274
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		if data[p] == 238 {
			goto st383
		}
		goto st273
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		if data[p] == 219 {
			goto st731
		}
		goto st5
	st731:
		if p++; p == pe {
			goto _test_eof731
		}
	st_case_731:
		goto tr774
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		if data[p] == 55 {
			goto st385
		}
		goto st274
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		if data[p] == 122 {
			goto st386
		}
		goto st273
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		if data[p] == 88 {
			goto st387
		}
		goto st5
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
		if data[p] == 90 {
			goto st388
		}
		goto st6
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
		if data[p] == 0 {
			goto st732
		}
		goto st7
	st732:
		if p++; p == pe {
			goto _test_eof732
		}
	st_case_732:
		goto tr775
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		switch data[p] {
		case 216:
			goto st390
		case 241:
			goto st734
		case 249:
			goto st735
		case 251:
			goto st736
		}
		goto st274
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
		if data[p] == 255 {
			goto st733
		}
		goto st273
	st733:
		if p++; p == pe {
			goto _test_eof733
		}
	st_case_733:
		goto tr776
	st734:
		if p++; p == pe {
			goto _test_eof734
		}
	st_case_734:
		goto st656
	st735:
		if p++; p == pe {
			goto _test_eof735
		}
	st_case_735:
		goto tr777
	st736:
		if p++; p == pe {
			goto _test_eof736
		}
	st_case_736:
		goto tr778
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof149: cs = 149; goto _test_eof
	_test_eof150: cs = 150; goto _test_eof
	_test_eof151: cs = 151; goto _test_eof
	_test_eof152: cs = 152; goto _test_eof
	_test_eof153: cs = 153; goto _test_eof
	_test_eof154: cs = 154; goto _test_eof
	_test_eof155: cs = 155; goto _test_eof
	_test_eof156: cs = 156; goto _test_eof
	_test_eof157: cs = 157; goto _test_eof
	_test_eof158: cs = 158; goto _test_eof
	_test_eof159: cs = 159; goto _test_eof
	_test_eof160: cs = 160; goto _test_eof
	_test_eof161: cs = 161; goto _test_eof
	_test_eof162: cs = 162; goto _test_eof
	_test_eof163: cs = 163; goto _test_eof
	_test_eof164: cs = 164; goto _test_eof
	_test_eof165: cs = 165; goto _test_eof
	_test_eof166: cs = 166; goto _test_eof
	_test_eof167: cs = 167; goto _test_eof
	_test_eof168: cs = 168; goto _test_eof
	_test_eof169: cs = 169; goto _test_eof
	_test_eof170: cs = 170; goto _test_eof
	_test_eof171: cs = 171; goto _test_eof
	_test_eof172: cs = 172; goto _test_eof
	_test_eof173: cs = 173; goto _test_eof
	_test_eof174: cs = 174; goto _test_eof
	_test_eof175: cs = 175; goto _test_eof
	_test_eof176: cs = 176; goto _test_eof
	_test_eof177: cs = 177; goto _test_eof
	_test_eof178: cs = 178; goto _test_eof
	_test_eof179: cs = 179; goto _test_eof
	_test_eof180: cs = 180; goto _test_eof
	_test_eof181: cs = 181; goto _test_eof
	_test_eof182: cs = 182; goto _test_eof
	_test_eof183: cs = 183; goto _test_eof
	_test_eof184: cs = 184; goto _test_eof
	_test_eof185: cs = 185; goto _test_eof
	_test_eof186: cs = 186; goto _test_eof
	_test_eof187: cs = 187; goto _test_eof
	_test_eof188: cs = 188; goto _test_eof
	_test_eof189: cs = 189; goto _test_eof
	_test_eof190: cs = 190; goto _test_eof
	_test_eof191: cs = 191; goto _test_eof
	_test_eof192: cs = 192; goto _test_eof
	_test_eof193: cs = 193; goto _test_eof
	_test_eof194: cs = 194; goto _test_eof
	_test_eof195: cs = 195; goto _test_eof
	_test_eof196: cs = 196; goto _test_eof
	_test_eof197: cs = 197; goto _test_eof
	_test_eof198: cs = 198; goto _test_eof
	_test_eof199: cs = 199; goto _test_eof
	_test_eof200: cs = 200; goto _test_eof
	_test_eof201: cs = 201; goto _test_eof
	_test_eof202: cs = 202; goto _test_eof
	_test_eof203: cs = 203; goto _test_eof
	_test_eof204: cs = 204; goto _test_eof
	_test_eof205: cs = 205; goto _test_eof
	_test_eof206: cs = 206; goto _test_eof
	_test_eof207: cs = 207; goto _test_eof
	_test_eof208: cs = 208; goto _test_eof
	_test_eof209: cs = 209; goto _test_eof
	_test_eof210: cs = 210; goto _test_eof
	_test_eof211: cs = 211; goto _test_eof
	_test_eof212: cs = 212; goto _test_eof
	_test_eof213: cs = 213; goto _test_eof
	_test_eof214: cs = 214; goto _test_eof
	_test_eof215: cs = 215; goto _test_eof
	_test_eof216: cs = 216; goto _test_eof
	_test_eof217: cs = 217; goto _test_eof
	_test_eof218: cs = 218; goto _test_eof
	_test_eof219: cs = 219; goto _test_eof
	_test_eof220: cs = 220; goto _test_eof
	_test_eof221: cs = 221; goto _test_eof
	_test_eof222: cs = 222; goto _test_eof
	_test_eof223: cs = 223; goto _test_eof
	_test_eof224: cs = 224; goto _test_eof
	_test_eof225: cs = 225; goto _test_eof
	_test_eof226: cs = 226; goto _test_eof
	_test_eof227: cs = 227; goto _test_eof
	_test_eof228: cs = 228; goto _test_eof
	_test_eof229: cs = 229; goto _test_eof
	_test_eof230: cs = 230; goto _test_eof
	_test_eof231: cs = 231; goto _test_eof
	_test_eof232: cs = 232; goto _test_eof
	_test_eof233: cs = 233; goto _test_eof
	_test_eof234: cs = 234; goto _test_eof
	_test_eof235: cs = 235; goto _test_eof
	_test_eof236: cs = 236; goto _test_eof
	_test_eof237: cs = 237; goto _test_eof
	_test_eof238: cs = 238; goto _test_eof
	_test_eof239: cs = 239; goto _test_eof
	_test_eof240: cs = 240; goto _test_eof
	_test_eof241: cs = 241; goto _test_eof
	_test_eof242: cs = 242; goto _test_eof
	_test_eof243: cs = 243; goto _test_eof
	_test_eof244: cs = 244; goto _test_eof
	_test_eof245: cs = 245; goto _test_eof
	_test_eof246: cs = 246; goto _test_eof
	_test_eof247: cs = 247; goto _test_eof
	_test_eof248: cs = 248; goto _test_eof
	_test_eof249: cs = 249; goto _test_eof
	_test_eof250: cs = 250; goto _test_eof
	_test_eof251: cs = 251; goto _test_eof
	_test_eof252: cs = 252; goto _test_eof
	_test_eof253: cs = 253; goto _test_eof
	_test_eof254: cs = 254; goto _test_eof
	_test_eof255: cs = 255; goto _test_eof
	_test_eof256: cs = 256; goto _test_eof
	_test_eof257: cs = 257; goto _test_eof
	_test_eof258: cs = 258; goto _test_eof
	_test_eof259: cs = 259; goto _test_eof
	_test_eof260: cs = 260; goto _test_eof
	_test_eof261: cs = 261; goto _test_eof
	_test_eof262: cs = 262; goto _test_eof
	_test_eof391: cs = 391; goto _test_eof
	_test_eof392: cs = 392; goto _test_eof
	_test_eof263: cs = 263; goto _test_eof
	_test_eof264: cs = 264; goto _test_eof
	_test_eof265: cs = 265; goto _test_eof
	_test_eof266: cs = 266; goto _test_eof
	_test_eof267: cs = 267; goto _test_eof
	_test_eof268: cs = 268; goto _test_eof
	_test_eof269: cs = 269; goto _test_eof
	_test_eof270: cs = 270; goto _test_eof
	_test_eof271: cs = 271; goto _test_eof
	_test_eof393: cs = 393; goto _test_eof
	_test_eof394: cs = 394; goto _test_eof
	_test_eof395: cs = 395; goto _test_eof
	_test_eof396: cs = 396; goto _test_eof
	_test_eof397: cs = 397; goto _test_eof
	_test_eof398: cs = 398; goto _test_eof
	_test_eof399: cs = 399; goto _test_eof
	_test_eof400: cs = 400; goto _test_eof
	_test_eof401: cs = 401; goto _test_eof
	_test_eof402: cs = 402; goto _test_eof
	_test_eof403: cs = 403; goto _test_eof
	_test_eof404: cs = 404; goto _test_eof
	_test_eof405: cs = 405; goto _test_eof
	_test_eof406: cs = 406; goto _test_eof
	_test_eof407: cs = 407; goto _test_eof
	_test_eof408: cs = 408; goto _test_eof
	_test_eof409: cs = 409; goto _test_eof
	_test_eof410: cs = 410; goto _test_eof
	_test_eof411: cs = 411; goto _test_eof
	_test_eof412: cs = 412; goto _test_eof
	_test_eof413: cs = 413; goto _test_eof
	_test_eof414: cs = 414; goto _test_eof
	_test_eof415: cs = 415; goto _test_eof
	_test_eof416: cs = 416; goto _test_eof
	_test_eof417: cs = 417; goto _test_eof
	_test_eof418: cs = 418; goto _test_eof
	_test_eof419: cs = 419; goto _test_eof
	_test_eof420: cs = 420; goto _test_eof
	_test_eof421: cs = 421; goto _test_eof
	_test_eof422: cs = 422; goto _test_eof
	_test_eof423: cs = 423; goto _test_eof
	_test_eof424: cs = 424; goto _test_eof
	_test_eof425: cs = 425; goto _test_eof
	_test_eof426: cs = 426; goto _test_eof
	_test_eof427: cs = 427; goto _test_eof
	_test_eof428: cs = 428; goto _test_eof
	_test_eof429: cs = 429; goto _test_eof
	_test_eof430: cs = 430; goto _test_eof
	_test_eof431: cs = 431; goto _test_eof
	_test_eof432: cs = 432; goto _test_eof
	_test_eof433: cs = 433; goto _test_eof
	_test_eof434: cs = 434; goto _test_eof
	_test_eof435: cs = 435; goto _test_eof
	_test_eof436: cs = 436; goto _test_eof
	_test_eof437: cs = 437; goto _test_eof
	_test_eof438: cs = 438; goto _test_eof
	_test_eof439: cs = 439; goto _test_eof
	_test_eof440: cs = 440; goto _test_eof
	_test_eof441: cs = 441; goto _test_eof
	_test_eof442: cs = 442; goto _test_eof
	_test_eof443: cs = 443; goto _test_eof
	_test_eof444: cs = 444; goto _test_eof
	_test_eof445: cs = 445; goto _test_eof
	_test_eof446: cs = 446; goto _test_eof
	_test_eof447: cs = 447; goto _test_eof
	_test_eof448: cs = 448; goto _test_eof
	_test_eof449: cs = 449; goto _test_eof
	_test_eof450: cs = 450; goto _test_eof
	_test_eof451: cs = 451; goto _test_eof
	_test_eof452: cs = 452; goto _test_eof
	_test_eof453: cs = 453; goto _test_eof
	_test_eof454: cs = 454; goto _test_eof
	_test_eof455: cs = 455; goto _test_eof
	_test_eof456: cs = 456; goto _test_eof
	_test_eof457: cs = 457; goto _test_eof
	_test_eof458: cs = 458; goto _test_eof
	_test_eof459: cs = 459; goto _test_eof
	_test_eof460: cs = 460; goto _test_eof
	_test_eof461: cs = 461; goto _test_eof
	_test_eof462: cs = 462; goto _test_eof
	_test_eof463: cs = 463; goto _test_eof
	_test_eof464: cs = 464; goto _test_eof
	_test_eof465: cs = 465; goto _test_eof
	_test_eof466: cs = 466; goto _test_eof
	_test_eof467: cs = 467; goto _test_eof
	_test_eof468: cs = 468; goto _test_eof
	_test_eof469: cs = 469; goto _test_eof
	_test_eof470: cs = 470; goto _test_eof
	_test_eof471: cs = 471; goto _test_eof
	_test_eof472: cs = 472; goto _test_eof
	_test_eof473: cs = 473; goto _test_eof
	_test_eof474: cs = 474; goto _test_eof
	_test_eof475: cs = 475; goto _test_eof
	_test_eof476: cs = 476; goto _test_eof
	_test_eof477: cs = 477; goto _test_eof
	_test_eof478: cs = 478; goto _test_eof
	_test_eof479: cs = 479; goto _test_eof
	_test_eof480: cs = 480; goto _test_eof
	_test_eof481: cs = 481; goto _test_eof
	_test_eof482: cs = 482; goto _test_eof
	_test_eof483: cs = 483; goto _test_eof
	_test_eof484: cs = 484; goto _test_eof
	_test_eof485: cs = 485; goto _test_eof
	_test_eof486: cs = 486; goto _test_eof
	_test_eof487: cs = 487; goto _test_eof
	_test_eof488: cs = 488; goto _test_eof
	_test_eof489: cs = 489; goto _test_eof
	_test_eof490: cs = 490; goto _test_eof
	_test_eof491: cs = 491; goto _test_eof
	_test_eof492: cs = 492; goto _test_eof
	_test_eof493: cs = 493; goto _test_eof
	_test_eof494: cs = 494; goto _test_eof
	_test_eof495: cs = 495; goto _test_eof
	_test_eof496: cs = 496; goto _test_eof
	_test_eof497: cs = 497; goto _test_eof
	_test_eof498: cs = 498; goto _test_eof
	_test_eof499: cs = 499; goto _test_eof
	_test_eof500: cs = 500; goto _test_eof
	_test_eof501: cs = 501; goto _test_eof
	_test_eof502: cs = 502; goto _test_eof
	_test_eof503: cs = 503; goto _test_eof
	_test_eof504: cs = 504; goto _test_eof
	_test_eof505: cs = 505; goto _test_eof
	_test_eof506: cs = 506; goto _test_eof
	_test_eof507: cs = 507; goto _test_eof
	_test_eof508: cs = 508; goto _test_eof
	_test_eof509: cs = 509; goto _test_eof
	_test_eof510: cs = 510; goto _test_eof
	_test_eof511: cs = 511; goto _test_eof
	_test_eof512: cs = 512; goto _test_eof
	_test_eof513: cs = 513; goto _test_eof
	_test_eof514: cs = 514; goto _test_eof
	_test_eof515: cs = 515; goto _test_eof
	_test_eof516: cs = 516; goto _test_eof
	_test_eof517: cs = 517; goto _test_eof
	_test_eof518: cs = 518; goto _test_eof
	_test_eof519: cs = 519; goto _test_eof
	_test_eof520: cs = 520; goto _test_eof
	_test_eof521: cs = 521; goto _test_eof
	_test_eof522: cs = 522; goto _test_eof
	_test_eof523: cs = 523; goto _test_eof
	_test_eof524: cs = 524; goto _test_eof
	_test_eof525: cs = 525; goto _test_eof
	_test_eof526: cs = 526; goto _test_eof
	_test_eof527: cs = 527; goto _test_eof
	_test_eof528: cs = 528; goto _test_eof
	_test_eof529: cs = 529; goto _test_eof
	_test_eof530: cs = 530; goto _test_eof
	_test_eof531: cs = 531; goto _test_eof
	_test_eof532: cs = 532; goto _test_eof
	_test_eof533: cs = 533; goto _test_eof
	_test_eof534: cs = 534; goto _test_eof
	_test_eof535: cs = 535; goto _test_eof
	_test_eof536: cs = 536; goto _test_eof
	_test_eof537: cs = 537; goto _test_eof
	_test_eof538: cs = 538; goto _test_eof
	_test_eof539: cs = 539; goto _test_eof
	_test_eof540: cs = 540; goto _test_eof
	_test_eof541: cs = 541; goto _test_eof
	_test_eof542: cs = 542; goto _test_eof
	_test_eof543: cs = 543; goto _test_eof
	_test_eof544: cs = 544; goto _test_eof
	_test_eof545: cs = 545; goto _test_eof
	_test_eof546: cs = 546; goto _test_eof
	_test_eof547: cs = 547; goto _test_eof
	_test_eof548: cs = 548; goto _test_eof
	_test_eof549: cs = 549; goto _test_eof
	_test_eof550: cs = 550; goto _test_eof
	_test_eof551: cs = 551; goto _test_eof
	_test_eof552: cs = 552; goto _test_eof
	_test_eof553: cs = 553; goto _test_eof
	_test_eof554: cs = 554; goto _test_eof
	_test_eof555: cs = 555; goto _test_eof
	_test_eof556: cs = 556; goto _test_eof
	_test_eof557: cs = 557; goto _test_eof
	_test_eof558: cs = 558; goto _test_eof
	_test_eof559: cs = 559; goto _test_eof
	_test_eof560: cs = 560; goto _test_eof
	_test_eof561: cs = 561; goto _test_eof
	_test_eof562: cs = 562; goto _test_eof
	_test_eof563: cs = 563; goto _test_eof
	_test_eof564: cs = 564; goto _test_eof
	_test_eof565: cs = 565; goto _test_eof
	_test_eof566: cs = 566; goto _test_eof
	_test_eof567: cs = 567; goto _test_eof
	_test_eof568: cs = 568; goto _test_eof
	_test_eof569: cs = 569; goto _test_eof
	_test_eof570: cs = 570; goto _test_eof
	_test_eof571: cs = 571; goto _test_eof
	_test_eof572: cs = 572; goto _test_eof
	_test_eof573: cs = 573; goto _test_eof
	_test_eof574: cs = 574; goto _test_eof
	_test_eof575: cs = 575; goto _test_eof
	_test_eof576: cs = 576; goto _test_eof
	_test_eof577: cs = 577; goto _test_eof
	_test_eof578: cs = 578; goto _test_eof
	_test_eof579: cs = 579; goto _test_eof
	_test_eof580: cs = 580; goto _test_eof
	_test_eof581: cs = 581; goto _test_eof
	_test_eof582: cs = 582; goto _test_eof
	_test_eof583: cs = 583; goto _test_eof
	_test_eof584: cs = 584; goto _test_eof
	_test_eof585: cs = 585; goto _test_eof
	_test_eof586: cs = 586; goto _test_eof
	_test_eof587: cs = 587; goto _test_eof
	_test_eof588: cs = 588; goto _test_eof
	_test_eof589: cs = 589; goto _test_eof
	_test_eof590: cs = 590; goto _test_eof
	_test_eof591: cs = 591; goto _test_eof
	_test_eof592: cs = 592; goto _test_eof
	_test_eof593: cs = 593; goto _test_eof
	_test_eof594: cs = 594; goto _test_eof
	_test_eof595: cs = 595; goto _test_eof
	_test_eof596: cs = 596; goto _test_eof
	_test_eof597: cs = 597; goto _test_eof
	_test_eof598: cs = 598; goto _test_eof
	_test_eof599: cs = 599; goto _test_eof
	_test_eof600: cs = 600; goto _test_eof
	_test_eof601: cs = 601; goto _test_eof
	_test_eof602: cs = 602; goto _test_eof
	_test_eof603: cs = 603; goto _test_eof
	_test_eof604: cs = 604; goto _test_eof
	_test_eof605: cs = 605; goto _test_eof
	_test_eof606: cs = 606; goto _test_eof
	_test_eof607: cs = 607; goto _test_eof
	_test_eof608: cs = 608; goto _test_eof
	_test_eof609: cs = 609; goto _test_eof
	_test_eof610: cs = 610; goto _test_eof
	_test_eof611: cs = 611; goto _test_eof
	_test_eof612: cs = 612; goto _test_eof
	_test_eof613: cs = 613; goto _test_eof
	_test_eof614: cs = 614; goto _test_eof
	_test_eof615: cs = 615; goto _test_eof
	_test_eof616: cs = 616; goto _test_eof
	_test_eof617: cs = 617; goto _test_eof
	_test_eof618: cs = 618; goto _test_eof
	_test_eof619: cs = 619; goto _test_eof
	_test_eof620: cs = 620; goto _test_eof
	_test_eof621: cs = 621; goto _test_eof
	_test_eof622: cs = 622; goto _test_eof
	_test_eof623: cs = 623; goto _test_eof
	_test_eof624: cs = 624; goto _test_eof
	_test_eof625: cs = 625; goto _test_eof
	_test_eof626: cs = 626; goto _test_eof
	_test_eof627: cs = 627; goto _test_eof
	_test_eof628: cs = 628; goto _test_eof
	_test_eof629: cs = 629; goto _test_eof
	_test_eof630: cs = 630; goto _test_eof
	_test_eof631: cs = 631; goto _test_eof
	_test_eof632: cs = 632; goto _test_eof
	_test_eof633: cs = 633; goto _test_eof
	_test_eof634: cs = 634; goto _test_eof
	_test_eof635: cs = 635; goto _test_eof
	_test_eof636: cs = 636; goto _test_eof
	_test_eof637: cs = 637; goto _test_eof
	_test_eof638: cs = 638; goto _test_eof
	_test_eof639: cs = 639; goto _test_eof
	_test_eof640: cs = 640; goto _test_eof
	_test_eof641: cs = 641; goto _test_eof
	_test_eof272: cs = 272; goto _test_eof
	_test_eof642: cs = 642; goto _test_eof
	_test_eof643: cs = 643; goto _test_eof
	_test_eof644: cs = 644; goto _test_eof
	_test_eof645: cs = 645; goto _test_eof
	_test_eof646: cs = 646; goto _test_eof
	_test_eof647: cs = 647; goto _test_eof
	_test_eof648: cs = 648; goto _test_eof
	_test_eof649: cs = 649; goto _test_eof
	_test_eof650: cs = 650; goto _test_eof
	_test_eof651: cs = 651; goto _test_eof
	_test_eof273: cs = 273; goto _test_eof
	_test_eof274: cs = 274; goto _test_eof
	_test_eof275: cs = 275; goto _test_eof
	_test_eof276: cs = 276; goto _test_eof
	_test_eof277: cs = 277; goto _test_eof
	_test_eof278: cs = 278; goto _test_eof
	_test_eof279: cs = 279; goto _test_eof
	_test_eof280: cs = 280; goto _test_eof
	_test_eof652: cs = 652; goto _test_eof
	_test_eof281: cs = 281; goto _test_eof
	_test_eof282: cs = 282; goto _test_eof
	_test_eof283: cs = 283; goto _test_eof
	_test_eof653: cs = 653; goto _test_eof
	_test_eof654: cs = 654; goto _test_eof
	_test_eof655: cs = 655; goto _test_eof
	_test_eof656: cs = 656; goto _test_eof
	_test_eof284: cs = 284; goto _test_eof
	_test_eof285: cs = 285; goto _test_eof
	_test_eof286: cs = 286; goto _test_eof
	_test_eof287: cs = 287; goto _test_eof
	_test_eof288: cs = 288; goto _test_eof
	_test_eof289: cs = 289; goto _test_eof
	_test_eof657: cs = 657; goto _test_eof
	_test_eof658: cs = 658; goto _test_eof
	_test_eof659: cs = 659; goto _test_eof
	_test_eof660: cs = 660; goto _test_eof
	_test_eof661: cs = 661; goto _test_eof
	_test_eof662: cs = 662; goto _test_eof
	_test_eof663: cs = 663; goto _test_eof
	_test_eof664: cs = 664; goto _test_eof
	_test_eof665: cs = 665; goto _test_eof
	_test_eof666: cs = 666; goto _test_eof
	_test_eof667: cs = 667; goto _test_eof
	_test_eof668: cs = 668; goto _test_eof
	_test_eof669: cs = 669; goto _test_eof
	_test_eof670: cs = 670; goto _test_eof
	_test_eof671: cs = 671; goto _test_eof
	_test_eof290: cs = 290; goto _test_eof
	_test_eof291: cs = 291; goto _test_eof
	_test_eof292: cs = 292; goto _test_eof
	_test_eof293: cs = 293; goto _test_eof
	_test_eof294: cs = 294; goto _test_eof
	_test_eof672: cs = 672; goto _test_eof
	_test_eof295: cs = 295; goto _test_eof
	_test_eof673: cs = 673; goto _test_eof
	_test_eof296: cs = 296; goto _test_eof
	_test_eof297: cs = 297; goto _test_eof
	_test_eof674: cs = 674; goto _test_eof
	_test_eof298: cs = 298; goto _test_eof
	_test_eof299: cs = 299; goto _test_eof
	_test_eof300: cs = 300; goto _test_eof
	_test_eof301: cs = 301; goto _test_eof
	_test_eof302: cs = 302; goto _test_eof
	_test_eof675: cs = 675; goto _test_eof
	_test_eof303: cs = 303; goto _test_eof
	_test_eof304: cs = 304; goto _test_eof
	_test_eof305: cs = 305; goto _test_eof
	_test_eof676: cs = 676; goto _test_eof
	_test_eof306: cs = 306; goto _test_eof
	_test_eof677: cs = 677; goto _test_eof
	_test_eof307: cs = 307; goto _test_eof
	_test_eof678: cs = 678; goto _test_eof
	_test_eof308: cs = 308; goto _test_eof
	_test_eof309: cs = 309; goto _test_eof
	_test_eof679: cs = 679; goto _test_eof
	_test_eof310: cs = 310; goto _test_eof
	_test_eof311: cs = 311; goto _test_eof
	_test_eof680: cs = 680; goto _test_eof
	_test_eof312: cs = 312; goto _test_eof
	_test_eof313: cs = 313; goto _test_eof
	_test_eof314: cs = 314; goto _test_eof
	_test_eof681: cs = 681; goto _test_eof
	_test_eof315: cs = 315; goto _test_eof
	_test_eof316: cs = 316; goto _test_eof
	_test_eof317: cs = 317; goto _test_eof
	_test_eof682: cs = 682; goto _test_eof
	_test_eof318: cs = 318; goto _test_eof
	_test_eof319: cs = 319; goto _test_eof
	_test_eof320: cs = 320; goto _test_eof
	_test_eof321: cs = 321; goto _test_eof
	_test_eof683: cs = 683; goto _test_eof
	_test_eof684: cs = 684; goto _test_eof
	_test_eof685: cs = 685; goto _test_eof
	_test_eof686: cs = 686; goto _test_eof
	_test_eof322: cs = 322; goto _test_eof
	_test_eof323: cs = 323; goto _test_eof
	_test_eof687: cs = 687; goto _test_eof
	_test_eof324: cs = 324; goto _test_eof
	_test_eof325: cs = 325; goto _test_eof
	_test_eof326: cs = 326; goto _test_eof
	_test_eof327: cs = 327; goto _test_eof
	_test_eof688: cs = 688; goto _test_eof
	_test_eof328: cs = 328; goto _test_eof
	_test_eof329: cs = 329; goto _test_eof
	_test_eof330: cs = 330; goto _test_eof
	_test_eof331: cs = 331; goto _test_eof
	_test_eof689: cs = 689; goto _test_eof
	_test_eof332: cs = 332; goto _test_eof
	_test_eof333: cs = 333; goto _test_eof
	_test_eof334: cs = 334; goto _test_eof
	_test_eof690: cs = 690; goto _test_eof
	_test_eof335: cs = 335; goto _test_eof
	_test_eof336: cs = 336; goto _test_eof
	_test_eof337: cs = 337; goto _test_eof
	_test_eof338: cs = 338; goto _test_eof
	_test_eof339: cs = 339; goto _test_eof
	_test_eof340: cs = 340; goto _test_eof
	_test_eof691: cs = 691; goto _test_eof
	_test_eof692: cs = 692; goto _test_eof
	_test_eof341: cs = 341; goto _test_eof
	_test_eof342: cs = 342; goto _test_eof
	_test_eof343: cs = 343; goto _test_eof
	_test_eof693: cs = 693; goto _test_eof
	_test_eof344: cs = 344; goto _test_eof
	_test_eof345: cs = 345; goto _test_eof
	_test_eof346: cs = 346; goto _test_eof
	_test_eof694: cs = 694; goto _test_eof
	_test_eof695: cs = 695; goto _test_eof
	_test_eof696: cs = 696; goto _test_eof
	_test_eof697: cs = 697; goto _test_eof
	_test_eof698: cs = 698; goto _test_eof
	_test_eof699: cs = 699; goto _test_eof
	_test_eof700: cs = 700; goto _test_eof
	_test_eof701: cs = 701; goto _test_eof
	_test_eof702: cs = 702; goto _test_eof
	_test_eof703: cs = 703; goto _test_eof
	_test_eof704: cs = 704; goto _test_eof
	_test_eof705: cs = 705; goto _test_eof
	_test_eof706: cs = 706; goto _test_eof
	_test_eof707: cs = 707; goto _test_eof
	_test_eof708: cs = 708; goto _test_eof
	_test_eof709: cs = 709; goto _test_eof
	_test_eof710: cs = 710; goto _test_eof
	_test_eof711: cs = 711; goto _test_eof
	_test_eof712: cs = 712; goto _test_eof
	_test_eof713: cs = 713; goto _test_eof
	_test_eof714: cs = 714; goto _test_eof
	_test_eof715: cs = 715; goto _test_eof
	_test_eof716: cs = 716; goto _test_eof
	_test_eof717: cs = 717; goto _test_eof
	_test_eof718: cs = 718; goto _test_eof
	_test_eof719: cs = 719; goto _test_eof
	_test_eof720: cs = 720; goto _test_eof
	_test_eof721: cs = 721; goto _test_eof
	_test_eof722: cs = 722; goto _test_eof
	_test_eof723: cs = 723; goto _test_eof
	_test_eof347: cs = 347; goto _test_eof
	_test_eof348: cs = 348; goto _test_eof
	_test_eof349: cs = 349; goto _test_eof
	_test_eof350: cs = 350; goto _test_eof
	_test_eof351: cs = 351; goto _test_eof
	_test_eof352: cs = 352; goto _test_eof
	_test_eof353: cs = 353; goto _test_eof
	_test_eof354: cs = 354; goto _test_eof
	_test_eof724: cs = 724; goto _test_eof
	_test_eof355: cs = 355; goto _test_eof
	_test_eof356: cs = 356; goto _test_eof
	_test_eof357: cs = 357; goto _test_eof
	_test_eof358: cs = 358; goto _test_eof
	_test_eof359: cs = 359; goto _test_eof
	_test_eof725: cs = 725; goto _test_eof
	_test_eof360: cs = 360; goto _test_eof
	_test_eof361: cs = 361; goto _test_eof
	_test_eof362: cs = 362; goto _test_eof
	_test_eof726: cs = 726; goto _test_eof
	_test_eof363: cs = 363; goto _test_eof
	_test_eof364: cs = 364; goto _test_eof
	_test_eof365: cs = 365; goto _test_eof
	_test_eof727: cs = 727; goto _test_eof
	_test_eof366: cs = 366; goto _test_eof
	_test_eof367: cs = 367; goto _test_eof
	_test_eof368: cs = 368; goto _test_eof
	_test_eof369: cs = 369; goto _test_eof
	_test_eof370: cs = 370; goto _test_eof
	_test_eof371: cs = 371; goto _test_eof
	_test_eof372: cs = 372; goto _test_eof
	_test_eof373: cs = 373; goto _test_eof
	_test_eof374: cs = 374; goto _test_eof
	_test_eof728: cs = 728; goto _test_eof
	_test_eof375: cs = 375; goto _test_eof
	_test_eof376: cs = 376; goto _test_eof
	_test_eof377: cs = 377; goto _test_eof
	_test_eof729: cs = 729; goto _test_eof
	_test_eof378: cs = 378; goto _test_eof
	_test_eof379: cs = 379; goto _test_eof
	_test_eof380: cs = 380; goto _test_eof
	_test_eof730: cs = 730; goto _test_eof
	_test_eof381: cs = 381; goto _test_eof
	_test_eof382: cs = 382; goto _test_eof
	_test_eof383: cs = 383; goto _test_eof
	_test_eof731: cs = 731; goto _test_eof
	_test_eof384: cs = 384; goto _test_eof
	_test_eof385: cs = 385; goto _test_eof
	_test_eof386: cs = 386; goto _test_eof
	_test_eof387: cs = 387; goto _test_eof
	_test_eof388: cs = 388; goto _test_eof
	_test_eof732: cs = 732; goto _test_eof
	_test_eof389: cs = 389; goto _test_eof
	_test_eof390: cs = 390; goto _test_eof
	_test_eof733: cs = 733; goto _test_eof
	_test_eof734: cs = 734; goto _test_eof
	_test_eof735: cs = 735; goto _test_eof
	_test_eof736: cs = 736; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 722:
//line magic.rl:66
 return Epub;
		case 694, 723:
//line magic.rl:67
return Zip;
		case 391:
//line magic.rl:68
return Tar;
		case 725:
//line magic.rl:69
return Rar;
		case 653:
//line magic.rl:70
return Gzip;
		case 678:
//line magic.rl:71
return Bz2;
		case 675:
//line magic.rl:72
return SevenZip;
		case 674:
//line magic.rl:73
return Pdf
		case 692:
//line magic.rl:74
return Exe
		case 679:
//line magic.rl:75
return Swf
		case 728:
//line magic.rl:76
return Rtf
		case 693:
//line magic.rl:77
return Nes
		case 680:
//line magic.rl:78
return Crs
		case 687:
//line magic.rl:79
return Cab
		case 688:
//line magic.rl:80
return Eot
		case 673:
//line magic.rl:81
return Ps
		case 732:
//line magic.rl:82
return Xz
		case 671:
//line magic.rl:83
return Deb
		case 657:
//line magic.rl:84
return Ar
		case 655:
//line magic.rl:85
return Z
		case 689:
//line magic.rl:86
return Lz
		case 731:
//line magic.rl:87
return Rpm
		case 729:
//line magic.rl:88
return Elf
		case 681:
//line magic.rl:89
return Dcm
		case 652:
//line magic.rl:92
return Wasm
		case 691:
//line magic.rl:95
return Midi
		case 736:
//line magic.rl:96
return Mp3
		case 690:
//line magic.rl:97
return M4a
		case 727:
//line magic.rl:98
return Ogg
		case 724:
//line magic.rl:99
return Wav
		case 672:
//line magic.rl:100
return Amr
		case 735:
//line magic.rl:101
return Aac
		case 733:
//line magic.rl:104
return Jpeg
		case 393:
//line magic.rl:105
return Jpeg2000
		case 730:
//line magic.rl:106
return Png
		case 682:
//line magic.rl:107
return Gif
		case 726:
//line magic.rl:108
return Webp
		case 685:
//line magic.rl:109
return CR2
		case 683:
//line magic.rl:110
return Tiff
		case 677:
//line magic.rl:111
return Bmp
		case 686:
//line magic.rl:112
return Jxr
		case 676:
//line magic.rl:113
return Psd
		case 642:
//line magic.rl:114
return Ico
//line magic.go:7721
		}
	}

	_out: {}
	}

//line magic.rl:129

  return Unknown
}