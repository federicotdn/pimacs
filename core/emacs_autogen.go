// Automatically generated with pimacs.tools.extract
// DO NOT MODIFY!
// Generated from GNU Emacs commit: 76938e1789ecf447507199b6f1b7c0d4f7924ea3, branch: master

package core

const emacsConstant_charset_arg_max_autogen = 17
const emacsConstant_coding_arg_max_autogen = 13

func (es *emacsStubs) boolVector_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("bool-vector") // Source file: alloc.c
}

func (es *emacsStubs) cons_autogen(car, cdr lispObject) (lispObject, error) {
	return es.stub("cons") // Source file: alloc.c
}

func (es *emacsStubs) garbageCollect_autogen() (lispObject, error) {
	return es.stub("garbage-collect") // Source file: alloc.c
}

func (es *emacsStubs) garbageCollectMaybe_autogen(factor lispObject) (lispObject, error) {
	return es.stub("garbage-collect-maybe") // Source file: alloc.c
}

func (es *emacsStubs) list_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("list") // Source file: alloc.c
}

func (es *emacsStubs) makeBoolVector_autogen(length, init lispObject) (lispObject, error) {
	return es.stub("make-bool-vector") // Source file: alloc.c
}

func (es *emacsStubs) makeByteCode_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("make-byte-code") // Source file: alloc.c
}

func (es *emacsStubs) makeClosure_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("make-closure") // Source file: alloc.c
}

func (es *emacsStubs) makeFinalizer_autogen(function lispObject) (lispObject, error) {
	return es.stub("make-finalizer") // Source file: alloc.c
}

func (es *emacsStubs) makeList_autogen(length, init lispObject) (lispObject, error) {
	return es.stub("make-list") // Source file: alloc.c
}

func (es *emacsStubs) makeMarker_autogen() (lispObject, error) {
	return es.stub("make-marker") // Source file: alloc.c
}

func (es *emacsStubs) makeRecord_autogen(type_, slots, init lispObject) (lispObject, error) {
	return es.stub("make-record") // Source file: alloc.c
}

func (es *emacsStubs) makeString_autogen(length, init, multibyte lispObject) (lispObject, error) {
	return es.stub("make-string") // Source file: alloc.c
}

func (es *emacsStubs) makeSymbol_autogen(name lispObject) (lispObject, error) {
	return es.stub("make-symbol") // Source file: alloc.c
}

func (es *emacsStubs) makeVector_autogen(length, init lispObject) (lispObject, error) {
	return es.stub("make-vector") // Source file: alloc.c
}

func (es *emacsStubs) mallocInfo_autogen() (lispObject, error) {
	return es.stub("malloc-info") // Source file: alloc.c
}

func (es *emacsStubs) mallocTrim_autogen(leave_padding lispObject) (lispObject, error) {
	return es.stub("malloc-trim") // Source file: alloc.c
}

func (es *emacsStubs) memoryInfo_autogen() (lispObject, error) {
	return es.stub("memory-info") // Source file: alloc.c
}

func (es *emacsStubs) memoryUseCounts_autogen() (lispObject, error) {
	return es.stub("memory-use-counts") // Source file: alloc.c
}

func (es *emacsStubs) purecopy_autogen(obj lispObject) (lispObject, error) {
	return es.stub("purecopy") // Source file: alloc.c
}

func (es *emacsStubs) record_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("record") // Source file: alloc.c
}

func (es *emacsStubs) suspiciousObject_autogen(obj lispObject) (lispObject, error) {
	return es.stub("suspicious-object") // Source file: alloc.c
}

func (es *emacsStubs) vector_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("vector") // Source file: alloc.c
}

func (es *emacsStubs) debugTimerCheck_autogen() (lispObject, error) {
	return es.stub("debug-timer-check") // Source file: atimer.c
}

func (es *emacsStubs) barfIfBufferReadOnly_autogen(position lispObject) (lispObject, error) {
	return es.stub("barf-if-buffer-read-only") // Source file: buffer.c
}

func (es *emacsStubs) bufferBaseBuffer_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("buffer-base-buffer") // Source file: buffer.c
}

func (es *emacsStubs) bufferCharsModifiedTick_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("buffer-chars-modified-tick") // Source file: buffer.c
}

func (es *emacsStubs) bufferEnableUndo_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("buffer-enable-undo") // Source file: buffer.c
}

func (es *emacsStubs) bufferFileName_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("buffer-file-name") // Source file: buffer.c
}

func (es *emacsStubs) bufferList_autogen(frame lispObject) (lispObject, error) {
	return es.stub("buffer-list") // Source file: buffer.c
}

func (es *emacsStubs) bufferLiveP_autogen(object lispObject) (lispObject, error) {
	return es.stub("buffer-live-p") // Source file: buffer.c
}

func (es *emacsStubs) bufferLocalValue_autogen(variable, buffer lispObject) (lispObject, error) {
	return es.stub("buffer-local-value") // Source file: buffer.c
}

func (es *emacsStubs) bufferLocalVariables_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("buffer-local-variables") // Source file: buffer.c
}

func (es *emacsStubs) bufferModifiedP_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("buffer-modified-p") // Source file: buffer.c
}

func (es *emacsStubs) bufferModifiedTick_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("buffer-modified-tick") // Source file: buffer.c
}

func (es *emacsStubs) bufferName_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("buffer-name") // Source file: buffer.c
}

func (es *emacsStubs) bufferSwapText_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("buffer-swap-text") // Source file: buffer.c
}

func (es *emacsStubs) buryBufferInternal_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("bury-buffer-internal") // Source file: buffer.c
}

func (es *emacsStubs) currentBuffer_autogen() (lispObject, error) {
	return es.stub("current-buffer") // Source file: buffer.c
}

func (es *emacsStubs) deleteAllOverlays_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("delete-all-overlays") // Source file: buffer.c
}

func (es *emacsStubs) deleteOverlay_autogen(overlay lispObject) (lispObject, error) {
	return es.stub("delete-overlay") // Source file: buffer.c
}

func (es *emacsStubs) eraseBuffer_autogen() (lispObject, error) {
	return es.stub("erase-buffer") // Source file: buffer.c
}

func (es *emacsStubs) forceModeLineUpdate_autogen(all lispObject) (lispObject, error) {
	return es.stub("force-mode-line-update") // Source file: buffer.c
}

func (es *emacsStubs) generateNewBufferName_autogen(name, ignore lispObject) (lispObject, error) {
	return es.stub("generate-new-buffer-name") // Source file: buffer.c
}

func (es *emacsStubs) getBuffer_autogen(buffer_or_name lispObject) (lispObject, error) {
	return es.stub("get-buffer") // Source file: buffer.c
}

func (es *emacsStubs) getBufferCreate_autogen(buffer_or_name, inhibit_buffer_hooks lispObject) (lispObject, error) {
	return es.stub("get-buffer-create") // Source file: buffer.c
}

func (es *emacsStubs) getFileBuffer_autogen(filename lispObject) (lispObject, error) {
	return es.stub("get-file-buffer") // Source file: buffer.c
}

func (es *emacsStubs) internalSetBufferModifiedTick_autogen(tick, buffer lispObject) (lispObject, error) {
	return es.stub("internal--set-buffer-modified-tick") // Source file: buffer.c
}

func (es *emacsStubs) killAllLocalVariables_autogen(kill_permanent lispObject) (lispObject, error) {
	return es.stub("kill-all-local-variables") // Source file: buffer.c
}

func (es *emacsStubs) killBuffer_autogen(buffer_or_name lispObject) (lispObject, error) {
	return es.stub("kill-buffer") // Source file: buffer.c
}

func (es *emacsStubs) makeIndirectBuffer_autogen(base_buffer, name, clone, inhibit_buffer_hooks lispObject) (lispObject, error) {
	return es.stub("make-indirect-buffer") // Source file: buffer.c
}

func (es *emacsStubs) makeOverlay_autogen(beg, end, buffer, front_advance, rear_advance lispObject) (lispObject, error) {
	return es.stub("make-overlay") // Source file: buffer.c
}

func (es *emacsStubs) moveOverlay_autogen(overlay, beg, end, buffer lispObject) (lispObject, error) {
	return es.stub("move-overlay") // Source file: buffer.c
}

func (es *emacsStubs) nextOverlayChange_autogen(pos lispObject) (lispObject, error) {
	return es.stub("next-overlay-change") // Source file: buffer.c
}

func (es *emacsStubs) otherBuffer_autogen(buffer, visible_ok, frame lispObject) (lispObject, error) {
	return es.stub("other-buffer") // Source file: buffer.c
}

func (es *emacsStubs) overlayBuffer_autogen(overlay lispObject) (lispObject, error) {
	return es.stub("overlay-buffer") // Source file: buffer.c
}

func (es *emacsStubs) overlayEnd_autogen(overlay lispObject) (lispObject, error) {
	return es.stub("overlay-end") // Source file: buffer.c
}

func (es *emacsStubs) overlayGet_autogen(overlay, prop lispObject) (lispObject, error) {
	return es.stub("overlay-get") // Source file: buffer.c
}

func (es *emacsStubs) overlayLists_autogen() (lispObject, error) {
	return es.stub("overlay-lists") // Source file: buffer.c
}

func (es *emacsStubs) overlayProperties_autogen(overlay lispObject) (lispObject, error) {
	return es.stub("overlay-properties") // Source file: buffer.c
}

func (es *emacsStubs) overlayPut_autogen(overlay, prop, value lispObject) (lispObject, error) {
	return es.stub("overlay-put") // Source file: buffer.c
}

func (es *emacsStubs) overlayRecenter_autogen(pos lispObject) (lispObject, error) {
	return es.stub("overlay-recenter") // Source file: buffer.c
}

func (es *emacsStubs) overlayStart_autogen(overlay lispObject) (lispObject, error) {
	return es.stub("overlay-start") // Source file: buffer.c
}

func (es *emacsStubs) overlayTree_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("overlay-tree") // Source file: buffer.c
}

func (es *emacsStubs) overlayp_autogen(object lispObject) (lispObject, error) {
	return es.stub("overlayp") // Source file: buffer.c
}

func (es *emacsStubs) overlaysAt_autogen(pos, sorted lispObject) (lispObject, error) {
	return es.stub("overlays-at") // Source file: buffer.c
}

func (es *emacsStubs) overlaysIn_autogen(beg, end lispObject) (lispObject, error) {
	return es.stub("overlays-in") // Source file: buffer.c
}

func (es *emacsStubs) previousOverlayChange_autogen(pos lispObject) (lispObject, error) {
	return es.stub("previous-overlay-change") // Source file: buffer.c
}

func (es *emacsStubs) renameBuffer_autogen(newname, unique lispObject) (lispObject, error) {
	return es.stub("rename-buffer") // Source file: buffer.c
}

func (es *emacsStubs) restoreBufferModifiedP_autogen(flag lispObject) (lispObject, error) {
	return es.stub("restore-buffer-modified-p") // Source file: buffer.c
}

func (es *emacsStubs) setBuffer_autogen(buffer_or_name lispObject) (lispObject, error) {
	return es.stub("set-buffer") // Source file: buffer.c
}

func (es *emacsStubs) setBufferMajorMode_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("set-buffer-major-mode") // Source file: buffer.c
}

func (es *emacsStubs) setBufferModifiedP_autogen(flag lispObject) (lispObject, error) {
	return es.stub("set-buffer-modified-p") // Source file: buffer.c
}

func (es *emacsStubs) setBufferMultibyte_autogen(flag lispObject) (lispObject, error) {
	return es.stub("set-buffer-multibyte") // Source file: buffer.c
}

func (es *emacsStubs) byteCode_autogen(bytestr, vector, maxdepth lispObject) (lispObject, error) {
	return es.stub("byte-code") // Source file: bytecode.c
}

func (es *emacsStubs) internalStackStats_autogen() (lispObject, error) {
	return es.stub("internal-stack-stats") // Source file: bytecode.c
}

func (es *emacsStubs) callInteractively_autogen(function, record_flag, keys lispObject) (lispObject, error) {
	return es.stub("call-interactively") // Source file: callint.c
}

func (es *emacsStubs) funcallInteractively_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("funcall-interactively") // Source file: callint.c
}

func (es *emacsStubs) interactive_autogen(args lispObject) (lispObject, error) {
	return es.stub("interactive") // Source file: callint.c, attributes: const
}

func (es *emacsStubs) prefixNumericValue_autogen(raw lispObject) (lispObject, error) {
	return es.stub("prefix-numeric-value") // Source file: callint.c
}

func (es *emacsStubs) callProcess_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("call-process") // Source file: callproc.c
}

func (es *emacsStubs) callProcessRegion_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("call-process-region") // Source file: callproc.c
}

func (es *emacsStubs) getenvInternal_autogen(variable, env lispObject) (lispObject, error) {
	return es.stub("getenv-internal") // Source file: callproc.c
}

func (es *emacsStubs) capitalize_autogen(obj lispObject) (lispObject, error) {
	return es.stub("capitalize") // Source file: casefiddle.c
}

func (es *emacsStubs) capitalizeRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return es.stub("capitalize-region") // Source file: casefiddle.c
}

func (es *emacsStubs) capitalizeWord_autogen(arg lispObject) (lispObject, error) {
	return es.stub("capitalize-word") // Source file: casefiddle.c
}

func (es *emacsStubs) downcase_autogen(obj lispObject) (lispObject, error) {
	return es.stub("downcase") // Source file: casefiddle.c
}

func (es *emacsStubs) downcaseRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return es.stub("downcase-region") // Source file: casefiddle.c
}

func (es *emacsStubs) downcaseWord_autogen(arg lispObject) (lispObject, error) {
	return es.stub("downcase-word") // Source file: casefiddle.c
}

func (es *emacsStubs) upcase_autogen(obj lispObject) (lispObject, error) {
	return es.stub("upcase") // Source file: casefiddle.c
}

func (es *emacsStubs) upcaseInitials_autogen(obj lispObject) (lispObject, error) {
	return es.stub("upcase-initials") // Source file: casefiddle.c
}

func (es *emacsStubs) upcaseInitialsRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return es.stub("upcase-initials-region") // Source file: casefiddle.c
}

func (es *emacsStubs) upcaseRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return es.stub("upcase-region") // Source file: casefiddle.c
}

func (es *emacsStubs) upcaseWord_autogen(arg lispObject) (lispObject, error) {
	return es.stub("upcase-word") // Source file: casefiddle.c
}

func (es *emacsStubs) caseTableP_autogen(object lispObject) (lispObject, error) {
	return es.stub("case-table-p") // Source file: casetab.c
}

func (es *emacsStubs) currentCaseTable_autogen() (lispObject, error) {
	return es.stub("current-case-table") // Source file: casetab.c
}

func (es *emacsStubs) setCaseTable_autogen(table lispObject) (lispObject, error) {
	return es.stub("set-case-table") // Source file: casetab.c
}

func (es *emacsStubs) setStandardCaseTable_autogen(table lispObject) (lispObject, error) {
	return es.stub("set-standard-case-table") // Source file: casetab.c
}

func (es *emacsStubs) standardCaseTable_autogen() (lispObject, error) {
	return es.stub("standard-case-table") // Source file: casetab.c
}

func (es *emacsStubs) categoryDocstring_autogen(category, table lispObject) (lispObject, error) {
	return es.stub("category-docstring") // Source file: category.c
}

func (es *emacsStubs) categorySetMnemonics_autogen(category_set lispObject) (lispObject, error) {
	return es.stub("category-set-mnemonics") // Source file: category.c
}

func (es *emacsStubs) categoryTable_autogen() (lispObject, error) {
	return es.stub("category-table") // Source file: category.c
}

func (es *emacsStubs) categoryTableP_autogen(arg lispObject) (lispObject, error) {
	return es.stub("category-table-p") // Source file: category.c
}

func (es *emacsStubs) charCategorySet_autogen(ch lispObject) (lispObject, error) {
	return es.stub("char-category-set") // Source file: category.c
}

func (es *emacsStubs) copyCategoryTable_autogen(table lispObject) (lispObject, error) {
	return es.stub("copy-category-table") // Source file: category.c
}

func (es *emacsStubs) defineCategory_autogen(category, docstring, table lispObject) (lispObject, error) {
	return es.stub("define-category") // Source file: category.c
}

func (es *emacsStubs) getUnusedCategory_autogen(table lispObject) (lispObject, error) {
	return es.stub("get-unused-category") // Source file: category.c
}

func (es *emacsStubs) makeCategorySet_autogen(categories lispObject) (lispObject, error) {
	return es.stub("make-category-set") // Source file: category.c
}

func (es *emacsStubs) makeCategoryTable_autogen() (lispObject, error) {
	return es.stub("make-category-table") // Source file: category.c
}

func (es *emacsStubs) modifyCategoryEntry_autogen(character, category, table, reset lispObject) (lispObject, error) {
	return es.stub("modify-category-entry") // Source file: category.c
}

func (es *emacsStubs) setCategoryTable_autogen(table lispObject) (lispObject, error) {
	return es.stub("set-category-table") // Source file: category.c
}

func (es *emacsStubs) standardCategoryTable_autogen() (lispObject, error) {
	return es.stub("standard-category-table") // Source file: category.c
}

func (es *emacsStubs) cclExecute_autogen(ccl_prog, reg lispObject) (lispObject, error) {
	return es.stub("ccl-execute") // Source file: ccl.c
}

func (es *emacsStubs) cclExecuteOnString_autogen(ccl_prog, status, str, contin, unibyte_p lispObject) (lispObject, error) {
	return es.stub("ccl-execute-on-string") // Source file: ccl.c
}

func (es *emacsStubs) cclProgramP_autogen(object lispObject) (lispObject, error) {
	return es.stub("ccl-program-p") // Source file: ccl.c
}

func (es *emacsStubs) registerCclProgram_autogen(name, ccl_prog lispObject) (lispObject, error) {
	return es.stub("register-ccl-program") // Source file: ccl.c
}

func (es *emacsStubs) registerCodeConversionMap_autogen(symbol, map_ lispObject) (lispObject, error) {
	return es.stub("register-code-conversion-map") // Source file: ccl.c
}

func (es *emacsStubs) charResolveModifiers_autogen(character lispObject) (lispObject, error) {
	return es.stub("char-resolve-modifiers") // Source file: character.c
}

func (es *emacsStubs) charWidth_autogen(ch lispObject) (lispObject, error) {
	return es.stub("char-width") // Source file: character.c
}

func (es *emacsStubs) characterp_autogen(object, ignore lispObject) (lispObject, error) {
	return es.stub("characterp") // Source file: character.c, attributes: const
}

func (es *emacsStubs) getByte_autogen(position, string lispObject) (lispObject, error) {
	return es.stub("get-byte") // Source file: character.c
}

func (es *emacsStubs) maxChar_autogen(unicode lispObject) (lispObject, error) {
	return es.stub("max-char") // Source file: character.c, attributes: const
}

func (es *emacsStubs) multibyteCharToUnibyte_autogen(ch lispObject) (lispObject, error) {
	return es.stub("multibyte-char-to-unibyte") // Source file: character.c
}

func (es *emacsStubs) string_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("string") // Source file: character.c
}

func (es *emacsStubs) stringWidth_autogen(str, from, to lispObject) (lispObject, error) {
	return es.stub("string-width") // Source file: character.c
}

func (es *emacsStubs) unibyteCharToMultibyte_autogen(ch lispObject) (lispObject, error) {
	return es.stub("unibyte-char-to-multibyte") // Source file: character.c
}

func (es *emacsStubs) unibyteString_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("unibyte-string") // Source file: character.c
}

func (es *emacsStubs) charCharset_autogen(ch, restriction lispObject) (lispObject, error) {
	return es.stub("char-charset") // Source file: charset.c
}

func (es *emacsStubs) charsetAfter_autogen(pos lispObject) (lispObject, error) {
	return es.stub("charset-after") // Source file: charset.c
}

func (es *emacsStubs) charsetIdInternal_autogen(charset lispObject) (lispObject, error) {
	return es.stub("charset-id-internal") // Source file: charset.c
}

func (es *emacsStubs) charsetPlist_autogen(charset lispObject) (lispObject, error) {
	return es.stub("charset-plist") // Source file: charset.c
}

func (es *emacsStubs) charsetPriorityList_autogen(highestp lispObject) (lispObject, error) {
	return es.stub("charset-priority-list") // Source file: charset.c
}

func (es *emacsStubs) charsetp_autogen(object lispObject) (lispObject, error) {
	return es.stub("charsetp") // Source file: charset.c
}

func (es *emacsStubs) clearCharsetMaps_autogen() (lispObject, error) {
	return es.stub("clear-charset-maps") // Source file: charset.c
}

func (es *emacsStubs) declareEquivCharset_autogen(dimension, chars, final_char, charset lispObject) (lispObject, error) {
	return es.stub("declare-equiv-charset") // Source file: charset.c
}

func (es *emacsStubs) decodeChar_autogen(charset, code_point lispObject) (lispObject, error) {
	return es.stub("decode-char") // Source file: charset.c
}

func (es *emacsStubs) defineCharsetAlias_autogen(alias, charset lispObject) (lispObject, error) {
	return es.stub("define-charset-alias") // Source file: charset.c
}

func (es *emacsStubs) defineCharsetInternal_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("define-charset-internal") // Source file: charset.c
}

func (es *emacsStubs) encodeChar_autogen(ch, charset lispObject) (lispObject, error) {
	return es.stub("encode-char") // Source file: charset.c
}

func (es *emacsStubs) findCharsetRegion_autogen(beg, end, table lispObject) (lispObject, error) {
	return es.stub("find-charset-region") // Source file: charset.c
}

func (es *emacsStubs) findCharsetString_autogen(str, table lispObject) (lispObject, error) {
	return es.stub("find-charset-string") // Source file: charset.c
}

func (es *emacsStubs) getUnusedIsoFinalChar_autogen(dimension, chars lispObject) (lispObject, error) {
	return es.stub("get-unused-iso-final-char") // Source file: charset.c
}

func (es *emacsStubs) isoCharset_autogen(dimension, chars, final_char lispObject) (lispObject, error) {
	return es.stub("iso-charset") // Source file: charset.c
}

func (es *emacsStubs) makeChar_autogen(charset, code1, code2, code3, code4 lispObject) (lispObject, error) {
	return es.stub("make-char") // Source file: charset.c
}

func (es *emacsStubs) mapCharsetChars_autogen(function, charset, arg, from_code, to_code lispObject) (lispObject, error) {
	return es.stub("map-charset-chars") // Source file: charset.c
}

func (es *emacsStubs) setCharsetPlist_autogen(charset, plist lispObject) (lispObject, error) {
	return es.stub("set-charset-plist") // Source file: charset.c
}

func (es *emacsStubs) setCharsetPriority_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("set-charset-priority") // Source file: charset.c
}

func (es *emacsStubs) sortCharsets_autogen(charsets lispObject) (lispObject, error) {
	return es.stub("sort-charsets") // Source file: charset.c
}

func (es *emacsStubs) splitChar_autogen(ch lispObject) (lispObject, error) {
	return es.stub("split-char") // Source file: charset.c
}

func (es *emacsStubs) unifyCharset_autogen(charset, unify_map, deunify lispObject) (lispObject, error) {
	return es.stub("unify-charset") // Source file: charset.c
}

func (es *emacsStubs) charTableExtraSlot_autogen(char_table, n lispObject) (lispObject, error) {
	return es.stub("char-table-extra-slot") // Source file: chartab.c
}

func (es *emacsStubs) charTableParent_autogen(char_table lispObject) (lispObject, error) {
	return es.stub("char-table-parent") // Source file: chartab.c
}

func (es *emacsStubs) charTableRange_autogen(char_table, range_ lispObject) (lispObject, error) {
	return es.stub("char-table-range") // Source file: chartab.c
}

func (es *emacsStubs) charTableSubtype_autogen(char_table lispObject) (lispObject, error) {
	return es.stub("char-table-subtype") // Source file: chartab.c
}

func (es *emacsStubs) getUnicodePropertyInternal_autogen(char_table, ch lispObject) (lispObject, error) {
	return es.stub("get-unicode-property-internal") // Source file: chartab.c
}

func (es *emacsStubs) makeCharTable_autogen(purpose, init lispObject) (lispObject, error) {
	return es.stub("make-char-table") // Source file: chartab.c
}

func (es *emacsStubs) mapCharTable_autogen(function, char_table lispObject) (lispObject, error) {
	return es.stub("map-char-table") // Source file: chartab.c
}

func (es *emacsStubs) optimizeCharTable_autogen(char_table, test lispObject) (lispObject, error) {
	return es.stub("optimize-char-table") // Source file: chartab.c
}

func (es *emacsStubs) putUnicodePropertyInternal_autogen(char_table, ch, value lispObject) (lispObject, error) {
	return es.stub("put-unicode-property-internal") // Source file: chartab.c
}

func (es *emacsStubs) setCharTableExtraSlot_autogen(char_table, n, value lispObject) (lispObject, error) {
	return es.stub("set-char-table-extra-slot") // Source file: chartab.c
}

func (es *emacsStubs) setCharTableParent_autogen(char_table, parent lispObject) (lispObject, error) {
	return es.stub("set-char-table-parent") // Source file: chartab.c
}

func (es *emacsStubs) setCharTableRange_autogen(char_table, range_, value lispObject) (lispObject, error) {
	return es.stub("set-char-table-range") // Source file: chartab.c
}

func (es *emacsStubs) unicodePropertyTableInternal_autogen(prop lispObject) (lispObject, error) {
	return es.stub("unicode-property-table-internal") // Source file: chartab.c
}

func (es *emacsStubs) backwardChar_autogen(n lispObject) (lispObject, error) {
	return es.stub("backward-char") // Source file: cmds.c
}

func (es *emacsStubs) beginningOfLine_autogen(n lispObject) (lispObject, error) {
	return es.stub("beginning-of-line") // Source file: cmds.c
}

func (es *emacsStubs) deleteChar_autogen(n, killflag lispObject) (lispObject, error) {
	return es.stub("delete-char") // Source file: cmds.c
}

func (es *emacsStubs) endOfLine_autogen(n lispObject) (lispObject, error) {
	return es.stub("end-of-line") // Source file: cmds.c
}

func (es *emacsStubs) forwardChar_autogen(n lispObject) (lispObject, error) {
	return es.stub("forward-char") // Source file: cmds.c
}

func (es *emacsStubs) forwardLine_autogen(n lispObject) (lispObject, error) {
	return es.stub("forward-line") // Source file: cmds.c
}

func (es *emacsStubs) selfInsertCommand_autogen(n, c lispObject) (lispObject, error) {
	return es.stub("self-insert-command") // Source file: cmds.c
}

func (es *emacsStubs) checkCodingSystem_autogen(coding_system lispObject) (lispObject, error) {
	return es.stub("check-coding-system") // Source file: coding.c
}

func (es *emacsStubs) checkCodingSystemsRegion_autogen(start, end, coding_system_list lispObject) (lispObject, error) {
	return es.stub("check-coding-systems-region") // Source file: coding.c
}

func (es *emacsStubs) codingSystemAliases_autogen(coding_system lispObject) (lispObject, error) {
	return es.stub("coding-system-aliases") // Source file: coding.c
}

func (es *emacsStubs) codingSystemBase_autogen(coding_system lispObject) (lispObject, error) {
	return es.stub("coding-system-base") // Source file: coding.c
}

func (es *emacsStubs) codingSystemEolType_autogen(coding_system lispObject) (lispObject, error) {
	return es.stub("coding-system-eol-type") // Source file: coding.c
}

func (es *emacsStubs) codingSystemP_autogen(object lispObject) (lispObject, error) {
	return es.stub("coding-system-p") // Source file: coding.c
}

func (es *emacsStubs) codingSystemPlist_autogen(coding_system lispObject) (lispObject, error) {
	return es.stub("coding-system-plist") // Source file: coding.c
}

func (es *emacsStubs) codingSystemPriorityList_autogen(highestp lispObject) (lispObject, error) {
	return es.stub("coding-system-priority-list") // Source file: coding.c
}

func (es *emacsStubs) codingSystemPut_autogen(coding_system, prop, val lispObject) (lispObject, error) {
	return es.stub("coding-system-put") // Source file: coding.c
}

func (es *emacsStubs) decodeBig5Char_autogen(code lispObject) (lispObject, error) {
	return es.stub("decode-big5-char") // Source file: coding.c
}

func (es *emacsStubs) decodeCodingRegion_autogen(start, end, coding_system, destination lispObject) (lispObject, error) {
	return es.stub("decode-coding-region") // Source file: coding.c
}

func (es *emacsStubs) decodeCodingString_autogen(string, coding_system, nocopy, buffer lispObject) (lispObject, error) {
	return es.stub("decode-coding-string") // Source file: coding.c
}

func (es *emacsStubs) decodeSjisChar_autogen(code lispObject) (lispObject, error) {
	return es.stub("decode-sjis-char") // Source file: coding.c
}

func (es *emacsStubs) defineCodingSystemAlias_autogen(alias, coding_system lispObject) (lispObject, error) {
	return es.stub("define-coding-system-alias") // Source file: coding.c
}

func (es *emacsStubs) defineCodingSystemInternal_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("define-coding-system-internal") // Source file: coding.c
}

func (es *emacsStubs) detectCodingRegion_autogen(start, end, highest lispObject) (lispObject, error) {
	return es.stub("detect-coding-region") // Source file: coding.c
}

func (es *emacsStubs) detectCodingString_autogen(string, highest lispObject) (lispObject, error) {
	return es.stub("detect-coding-string") // Source file: coding.c
}

func (es *emacsStubs) encodeBig5Char_autogen(ch lispObject) (lispObject, error) {
	return es.stub("encode-big5-char") // Source file: coding.c
}

func (es *emacsStubs) encodeCodingRegion_autogen(start, end, coding_system, destination lispObject) (lispObject, error) {
	return es.stub("encode-coding-region") // Source file: coding.c
}

func (es *emacsStubs) encodeCodingString_autogen(string, coding_system, nocopy, buffer lispObject) (lispObject, error) {
	return es.stub("encode-coding-string") // Source file: coding.c
}

func (es *emacsStubs) encodeSjisChar_autogen(ch lispObject) (lispObject, error) {
	return es.stub("encode-sjis-char") // Source file: coding.c
}

func (es *emacsStubs) findCodingSystemsRegionInternal_autogen(start, end, exclude lispObject) (lispObject, error) {
	return es.stub("find-coding-systems-region-internal") // Source file: coding.c
}

func (es *emacsStubs) findOperationCodingSystem_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("find-operation-coding-system") // Source file: coding.c
}

func (es *emacsStubs) internalDecodeStringUtf8_autogen(string, buffer, nocopy, handle_8_bit, handle_over_uni, decode_method, count lispObject) (lispObject, error) {
	return es.stub("internal-decode-string-utf-8") // Source file: coding.c
}

func (es *emacsStubs) internalEncodeStringUtf8_autogen(string, buffer, nocopy, handle_8_bit, handle_over_uni, encode_method, count lispObject) (lispObject, error) {
	return es.stub("internal-encode-string-utf-8") // Source file: coding.c
}

func (es *emacsStubs) keyboardCodingSystem_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("keyboard-coding-system") // Source file: coding.c
}

func (es *emacsStubs) readCodingSystem_autogen(prompt, default_coding_system lispObject) (lispObject, error) {
	return es.stub("read-coding-system") // Source file: coding.c
}

func (es *emacsStubs) readNonNilCodingSystem_autogen(prompt lispObject) (lispObject, error) {
	return es.stub("read-non-nil-coding-system") // Source file: coding.c
}

func (es *emacsStubs) setCodingSystemPriority_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("set-coding-system-priority") // Source file: coding.c
}

func (es *emacsStubs) setKeyboardCodingSystemInternal_autogen(coding_system, terminal lispObject) (lispObject, error) {
	return es.stub("set-keyboard-coding-system-internal") // Source file: coding.c
}

func (es *emacsStubs) setSafeTerminalCodingSystemInternal_autogen(coding_system lispObject) (lispObject, error) {
	return es.stub("set-safe-terminal-coding-system-internal") // Source file: coding.c
}

func (es *emacsStubs) setTerminalCodingSystemInternal_autogen(coding_system, terminal lispObject) (lispObject, error) {
	return es.stub("set-terminal-coding-system-internal") // Source file: coding.c
}

func (es *emacsStubs) terminalCodingSystem_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("terminal-coding-system") // Source file: coding.c
}

func (es *emacsStubs) unencodableCharPosition_autogen(start, end, coding_system, count, string lispObject) (lispObject, error) {
	return es.stub("unencodable-char-position") // Source file: coding.c
}

func (es *emacsStubs) compCompileCtxtToFile_autogen(filename lispObject) (lispObject, error) {
	return es.stub("comp--compile-ctxt-to-file") // Source file: comp.c
}

func (es *emacsStubs) compInitCtxt_autogen() (lispObject, error) {
	return es.stub("comp--init-ctxt") // Source file: comp.c
}

func (es *emacsStubs) compInstallTrampoline_autogen(subr_name, trampoline lispObject) (lispObject, error) {
	return es.stub("comp--install-trampoline") // Source file: comp.c
}

func (es *emacsStubs) compLateRegisterSubr_autogen(name, c_name, minarg, maxarg, type_, rest, comp_u lispObject) (lispObject, error) {
	return es.stub("comp--late-register-subr") // Source file: comp.c
}

func (es *emacsStubs) compRegisterLambda_autogen(reloc_idx, c_name, minarg, maxarg, type_, rest, comp_u lispObject) (lispObject, error) {
	return es.stub("comp--register-lambda") // Source file: comp.c
}

func (es *emacsStubs) compRegisterSubr_autogen(name, c_name, minarg, maxarg, type_, rest, comp_u lispObject) (lispObject, error) {
	return es.stub("comp--register-subr") // Source file: comp.c
}

func (es *emacsStubs) compReleaseCtxt_autogen() (lispObject, error) {
	return es.stub("comp--release-ctxt") // Source file: comp.c
}

func (es *emacsStubs) compSubrSignature_autogen(subr lispObject) (lispObject, error) {
	return es.stub("comp--subr-signature") // Source file: comp.c
}

func (es *emacsStubs) compElToElnFilename_autogen(filename, base_dir lispObject) (lispObject, error) {
	return es.stub("comp-el-to-eln-filename") // Source file: comp.c
}

func (es *emacsStubs) compElToElnRelFilename_autogen(filename lispObject) (lispObject, error) {
	return es.stub("comp-el-to-eln-rel-filename") // Source file: comp.c
}

func (es *emacsStubs) compLibgccjitVersion_autogen() (lispObject, error) {
	return es.stub("comp-libgccjit-version") // Source file: comp.c
}

func (es *emacsStubs) compNativeCompilerOptionsEffectiveP_autogen() (lispObject, error) {
	return es.stub("comp-native-compiler-options-effective-p") // Source file: comp.c
}

func (es *emacsStubs) compNativeDriverOptionsEffectiveP_autogen() (lispObject, error) {
	return es.stub("comp-native-driver-options-effective-p") // Source file: comp.c
}

func (es *emacsStubs) nativeCompAvailableP_autogen() (lispObject, error) {
	return es.stub("native-comp-available-p") // Source file: comp.c
}

func (es *emacsStubs) nativeElispLoad_autogen(filename, late_load lispObject) (lispObject, error) {
	return es.stub("native-elisp-load") // Source file: comp.c
}

func (es *emacsStubs) clearCompositionCache_autogen() (lispObject, error) {
	return es.stub("clear-composition-cache") // Source file: composite.c
}

func (es *emacsStubs) composeRegionInternal_autogen(start, end, components, modification_func lispObject) (lispObject, error) {
	return es.stub("compose-region-internal") // Source file: composite.c
}

func (es *emacsStubs) composeStringInternal_autogen(string, start, end, components, modification_func lispObject) (lispObject, error) {
	return es.stub("compose-string-internal") // Source file: composite.c
}

func (es *emacsStubs) compositionGetGstring_autogen(from, to, font_object, string lispObject) (lispObject, error) {
	return es.stub("composition-get-gstring") // Source file: composite.c
}

func (es *emacsStubs) compositionSortRules_autogen(rules lispObject) (lispObject, error) {
	return es.stub("composition-sort-rules") // Source file: composite.c
}

func (es *emacsStubs) findCompositionInternal_autogen(pos, limit, string, detail_p lispObject) (lispObject, error) {
	return es.stub("find-composition-internal") // Source file: composite.c
}

func (es *emacsStubs) cygwinConvertFileNameFromWindows_autogen(file, absolute_p lispObject) (lispObject, error) {
	return es.stub("cygwin-convert-file-name-from-windows") // Source file: cygw32.c
}

func (es *emacsStubs) cygwinConvertFileNameToWindows_autogen(file, absolute_p lispObject) (lispObject, error) {
	return es.stub("cygwin-convert-file-name-to-windows") // Source file: cygw32.c
}

func (es *emacsStubs) rem_autogen(x, y lispObject) (lispObject, error) {
	return es.stub("%") // Source file: data.c
}

func (es *emacsStubs) times_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("*") // Source file: data.c
}

func (es *emacsStubs) plus_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("+") // Source file: data.c
}

func (es *emacsStubs) minus_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("-") // Source file: data.c
}

func (es *emacsStubs) quo_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("/") // Source file: data.c
}

func (es *emacsStubs) neq_autogen(num1, num2 lispObject) (lispObject, error) {
	return es.stub("/=") // Source file: data.c
}

func (es *emacsStubs) add1_autogen(number lispObject) (lispObject, error) {
	return es.stub("1+") // Source file: data.c
}

func (es *emacsStubs) sub1_autogen(number lispObject) (lispObject, error) {
	return es.stub("1-") // Source file: data.c
}

func (es *emacsStubs) lss_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("<") // Source file: data.c
}

func (es *emacsStubs) leq_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("<=") // Source file: data.c
}

func (es *emacsStubs) eqlsign_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("=") // Source file: data.c
}

func (es *emacsStubs) gtr_autogen(args ...lispObject) (lispObject, error) {
	return es.stub(">") // Source file: data.c
}

func (es *emacsStubs) geq_autogen(args ...lispObject) (lispObject, error) {
	return es.stub(">=") // Source file: data.c
}

func (es *emacsStubs) addVariableWatcher_autogen(symbol, watch_function lispObject) (lispObject, error) {
	return es.stub("add-variable-watcher") // Source file: data.c
}

func (es *emacsStubs) aref_autogen(array, idx lispObject) (lispObject, error) {
	return es.stub("aref") // Source file: data.c
}

func (es *emacsStubs) arrayp_autogen(object lispObject) (lispObject, error) {
	return es.stub("arrayp") // Source file: data.c
}

func (es *emacsStubs) aset_autogen(array, idx, newelt lispObject) (lispObject, error) {
	return es.stub("aset") // Source file: data.c
}

func (es *emacsStubs) ash_autogen(value, count lispObject) (lispObject, error) {
	return es.stub("ash") // Source file: data.c
}

func (es *emacsStubs) atom_autogen(object lispObject) (lispObject, error) {
	return es.stub("atom") // Source file: data.c, attributes: const
}

func (es *emacsStubs) bareSymbol_autogen(sym lispObject) (lispObject, error) {
	return es.stub("bare-symbol") // Source file: data.c
}

func (es *emacsStubs) bareSymbolP_autogen(object lispObject) (lispObject, error) {
	return es.stub("bare-symbol-p") // Source file: data.c, attributes: const
}

func (es *emacsStubs) boolVectorCountConsecutive_autogen(a, b, i lispObject) (lispObject, error) {
	return es.stub("bool-vector-count-consecutive") // Source file: data.c
}

func (es *emacsStubs) boolVectorCountPopulation_autogen(a lispObject) (lispObject, error) {
	return es.stub("bool-vector-count-population") // Source file: data.c
}

func (es *emacsStubs) boolVectorExclusiveOr_autogen(a, b, c lispObject) (lispObject, error) {
	return es.stub("bool-vector-exclusive-or") // Source file: data.c
}

func (es *emacsStubs) boolVectorIntersection_autogen(a, b, c lispObject) (lispObject, error) {
	return es.stub("bool-vector-intersection") // Source file: data.c
}

func (es *emacsStubs) boolVectorNot_autogen(a, b lispObject) (lispObject, error) {
	return es.stub("bool-vector-not") // Source file: data.c
}

func (es *emacsStubs) boolVectorP_autogen(object lispObject) (lispObject, error) {
	return es.stub("bool-vector-p") // Source file: data.c
}

func (es *emacsStubs) boolVectorSetDifference_autogen(a, b, c lispObject) (lispObject, error) {
	return es.stub("bool-vector-set-difference") // Source file: data.c
}

func (es *emacsStubs) boolVectorSubsetp_autogen(a, b lispObject) (lispObject, error) {
	return es.stub("bool-vector-subsetp") // Source file: data.c
}

func (es *emacsStubs) boolVectorUnion_autogen(a, b, c lispObject) (lispObject, error) {
	return es.stub("bool-vector-union") // Source file: data.c
}

func (es *emacsStubs) boundp_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("boundp") // Source file: data.c
}

func (es *emacsStubs) bufferp_autogen(object lispObject) (lispObject, error) {
	return es.stub("bufferp") // Source file: data.c
}

func (es *emacsStubs) byteCodeFunctionP_autogen(object lispObject) (lispObject, error) {
	return es.stub("byte-code-function-p") // Source file: data.c
}

func (es *emacsStubs) byteorder_autogen() (lispObject, error) {
	return es.stub("byteorder") // Source file: data.c, attributes: const
}

func (es *emacsStubs) car_autogen(list lispObject) (lispObject, error) {
	return es.stub("car") // Source file: data.c
}

func (es *emacsStubs) carSafe_autogen(object lispObject) (lispObject, error) {
	return es.stub("car-safe") // Source file: data.c
}

func (es *emacsStubs) cdr_autogen(list lispObject) (lispObject, error) {
	return es.stub("cdr") // Source file: data.c
}

func (es *emacsStubs) cdrSafe_autogen(object lispObject) (lispObject, error) {
	return es.stub("cdr-safe") // Source file: data.c
}

func (es *emacsStubs) charOrStringP_autogen(object lispObject) (lispObject, error) {
	return es.stub("char-or-string-p") // Source file: data.c, attributes: const
}

func (es *emacsStubs) charTableP_autogen(object lispObject) (lispObject, error) {
	return es.stub("char-table-p") // Source file: data.c
}

func (es *emacsStubs) commandModes_autogen(command lispObject) (lispObject, error) {
	return es.stub("command-modes") // Source file: data.c
}

func (es *emacsStubs) conditionVariableP_autogen(object lispObject) (lispObject, error) {
	return es.stub("condition-variable-p") // Source file: data.c
}

func (es *emacsStubs) consp_autogen(object lispObject) (lispObject, error) {
	return es.stub("consp") // Source file: data.c, attributes: const
}

func (es *emacsStubs) defalias_autogen(symbol, definition, docstring lispObject) (lispObject, error) {
	return es.stub("defalias") // Source file: data.c
}

func (es *emacsStubs) defaultBoundp_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("default-boundp") // Source file: data.c
}

func (es *emacsStubs) defaultValue_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("default-value") // Source file: data.c
}

func (es *emacsStubs) eq_autogen(obj1, obj2 lispObject) (lispObject, error) {
	return es.stub("eq") // Source file: data.c, attributes: const
}

func (es *emacsStubs) fboundp_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("fboundp") // Source file: data.c
}

func (es *emacsStubs) floatp_autogen(object lispObject) (lispObject, error) {
	return es.stub("floatp") // Source file: data.c, attributes: const
}

func (es *emacsStubs) fmakunbound_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("fmakunbound") // Source file: data.c
}

func (es *emacsStubs) fset_autogen(symbol, definition lispObject) (lispObject, error) {
	return es.stub("fset") // Source file: data.c
}

func (es *emacsStubs) getVariableWatchers_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("get-variable-watchers") // Source file: data.c
}

func (es *emacsStubs) indirectFunction_autogen(object, noerror lispObject) (lispObject, error) {
	return es.stub("indirect-function") // Source file: data.c
}

func (es *emacsStubs) indirectVariable_autogen(object lispObject) (lispObject, error) {
	return es.stub("indirect-variable") // Source file: data.c
}

func (es *emacsStubs) integerOrMarkerP_autogen(object lispObject) (lispObject, error) {
	return es.stub("integer-or-marker-p") // Source file: data.c
}

func (es *emacsStubs) integerp_autogen(object lispObject) (lispObject, error) {
	return es.stub("integerp") // Source file: data.c, attributes: const
}

func (es *emacsStubs) interactiveForm_autogen(cmd lispObject) (lispObject, error) {
	return es.stub("interactive-form") // Source file: data.c
}

func (es *emacsStubs) keywordp_autogen(object lispObject) (lispObject, error) {
	return es.stub("keywordp") // Source file: data.c
}

func (es *emacsStubs) killLocalVariable_autogen(variable lispObject) (lispObject, error) {
	return es.stub("kill-local-variable") // Source file: data.c
}

func (es *emacsStubs) listp_autogen(object lispObject) (lispObject, error) {
	return es.stub("listp") // Source file: data.c, attributes: const
}

func (es *emacsStubs) localVariableIfSetP_autogen(variable, buffer lispObject) (lispObject, error) {
	return es.stub("local-variable-if-set-p") // Source file: data.c
}

func (es *emacsStubs) localVariableP_autogen(variable, buffer lispObject) (lispObject, error) {
	return es.stub("local-variable-p") // Source file: data.c
}

func (es *emacsStubs) logand_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("logand") // Source file: data.c
}

func (es *emacsStubs) logcount_autogen(value lispObject) (lispObject, error) {
	return es.stub("logcount") // Source file: data.c
}

func (es *emacsStubs) logior_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("logior") // Source file: data.c
}

func (es *emacsStubs) lognot_autogen(number lispObject) (lispObject, error) {
	return es.stub("lognot") // Source file: data.c
}

func (es *emacsStubs) logxor_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("logxor") // Source file: data.c
}

func (es *emacsStubs) makeLocalVariable_autogen(variable lispObject) (lispObject, error) {
	return es.stub("make-local-variable") // Source file: data.c
}

func (es *emacsStubs) makeVariableBufferLocal_autogen(variable lispObject) (lispObject, error) {
	return es.stub("make-variable-buffer-local") // Source file: data.c
}

func (es *emacsStubs) makunbound_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("makunbound") // Source file: data.c
}

func (es *emacsStubs) markerp_autogen(object lispObject) (lispObject, error) {
	return es.stub("markerp") // Source file: data.c
}

func (es *emacsStubs) max_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("max") // Source file: data.c
}

func (es *emacsStubs) min_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("min") // Source file: data.c
}

func (es *emacsStubs) mod_autogen(x, y lispObject) (lispObject, error) {
	return es.stub("mod") // Source file: data.c
}

func (es *emacsStubs) moduleFunctionP_autogen(object lispObject) (lispObject, error) {
	return es.stub("module-function-p") // Source file: data.c, attributes: const
}

func (es *emacsStubs) multibyteStringP_autogen(object lispObject) (lispObject, error) {
	return es.stub("multibyte-string-p") // Source file: data.c
}

func (es *emacsStubs) mutexp_autogen(object lispObject) (lispObject, error) {
	return es.stub("mutexp") // Source file: data.c
}

func (es *emacsStubs) nativeCompUnitFile_autogen(comp_unit lispObject) (lispObject, error) {
	return es.stub("native-comp-unit-file") // Source file: data.c
}

func (es *emacsStubs) nativeCompUnitSetFile_autogen(comp_unit, new_file lispObject) (lispObject, error) {
	return es.stub("native-comp-unit-set-file") // Source file: data.c
}

func (es *emacsStubs) natnump_autogen(object lispObject) (lispObject, error) {
	return es.stub("natnump") // Source file: data.c, attributes: const
}

func (es *emacsStubs) nlistp_autogen(object lispObject) (lispObject, error) {
	return es.stub("nlistp") // Source file: data.c, attributes: const
}

func (es *emacsStubs) null_autogen(object lispObject) (lispObject, error) {
	return es.stub("null") // Source file: data.c, attributes: const
}

func (es *emacsStubs) numberOrMarkerP_autogen(object lispObject) (lispObject, error) {
	return es.stub("number-or-marker-p") // Source file: data.c
}

func (es *emacsStubs) numberToString_autogen(number lispObject) (lispObject, error) {
	return es.stub("number-to-string") // Source file: data.c
}

func (es *emacsStubs) numberp_autogen(object lispObject) (lispObject, error) {
	return es.stub("numberp") // Source file: data.c, attributes: const
}

func (es *emacsStubs) positionSymbol_autogen(sym, pos lispObject) (lispObject, error) {
	return es.stub("position-symbol") // Source file: data.c
}

func (es *emacsStubs) recordp_autogen(object lispObject) (lispObject, error) {
	return es.stub("recordp") // Source file: data.c
}

func (es *emacsStubs) removePosFromSymbol_autogen(arg lispObject) (lispObject, error) {
	return es.stub("remove-pos-from-symbol") // Source file: data.c
}

func (es *emacsStubs) removeVariableWatcher_autogen(symbol, watch_function lispObject) (lispObject, error) {
	return es.stub("remove-variable-watcher") // Source file: data.c
}

func (es *emacsStubs) sequencep_autogen(object lispObject) (lispObject, error) {
	return es.stub("sequencep") // Source file: data.c
}

func (es *emacsStubs) set_autogen(symbol, newval lispObject) (lispObject, error) {
	return es.stub("set") // Source file: data.c
}

func (es *emacsStubs) setDefault_autogen(symbol, value lispObject) (lispObject, error) {
	return es.stub("set-default") // Source file: data.c
}

func (es *emacsStubs) setcar_autogen(cell, newcar lispObject) (lispObject, error) {
	return es.stub("setcar") // Source file: data.c
}

func (es *emacsStubs) setcdr_autogen(cell, newcdr lispObject) (lispObject, error) {
	return es.stub("setcdr") // Source file: data.c
}

func (es *emacsStubs) setplist_autogen(symbol, newplist lispObject) (lispObject, error) {
	return es.stub("setplist") // Source file: data.c
}

func (es *emacsStubs) stringToNumber_autogen(string, base lispObject) (lispObject, error) {
	return es.stub("string-to-number") // Source file: data.c
}

func (es *emacsStubs) stringp_autogen(object lispObject) (lispObject, error) {
	return es.stub("stringp") // Source file: data.c, attributes: const
}

func (es *emacsStubs) subrArity_autogen(subr lispObject) (lispObject, error) {
	return es.stub("subr-arity") // Source file: data.c
}

func (es *emacsStubs) subrName_autogen(subr lispObject) (lispObject, error) {
	return es.stub("subr-name") // Source file: data.c
}

func (es *emacsStubs) subrNativeCompUnit_autogen(subr lispObject) (lispObject, error) {
	return es.stub("subr-native-comp-unit") // Source file: data.c
}

func (es *emacsStubs) subrNativeElispP_autogen(object lispObject) (lispObject, error) {
	return es.stub("subr-native-elisp-p") // Source file: data.c
}

func (es *emacsStubs) subrNativeLambdaList_autogen(subr lispObject) (lispObject, error) {
	return es.stub("subr-native-lambda-list") // Source file: data.c
}

func (es *emacsStubs) subrType_autogen(subr lispObject) (lispObject, error) {
	return es.stub("subr-type") // Source file: data.c
}

func (es *emacsStubs) subrp_autogen(object lispObject) (lispObject, error) {
	return es.stub("subrp") // Source file: data.c
}

func (es *emacsStubs) symbolFunction_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("symbol-function") // Source file: data.c
}

func (es *emacsStubs) symbolName_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("symbol-name") // Source file: data.c
}

func (es *emacsStubs) symbolPlist_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("symbol-plist") // Source file: data.c
}

func (es *emacsStubs) symbolValue_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("symbol-value") // Source file: data.c
}

func (es *emacsStubs) symbolWithPosP_autogen(object lispObject) (lispObject, error) {
	return es.stub("symbol-with-pos-p") // Source file: data.c, attributes: const
}

func (es *emacsStubs) symbolWithPosPos_autogen(ls lispObject) (lispObject, error) {
	return es.stub("symbol-with-pos-pos") // Source file: data.c
}

func (es *emacsStubs) symbolp_autogen(object lispObject) (lispObject, error) {
	return es.stub("symbolp") // Source file: data.c, attributes: const
}

func (es *emacsStubs) threadp_autogen(object lispObject) (lispObject, error) {
	return es.stub("threadp") // Source file: data.c
}

func (es *emacsStubs) typeOf_autogen(object lispObject) (lispObject, error) {
	return es.stub("type-of") // Source file: data.c
}

func (es *emacsStubs) userPtrp_autogen(object lispObject) (lispObject, error) {
	return es.stub("user-ptrp") // Source file: data.c
}

func (es *emacsStubs) variableBindingLocus_autogen(variable lispObject) (lispObject, error) {
	return es.stub("variable-binding-locus") // Source file: data.c
}

func (es *emacsStubs) vectorOrCharTableP_autogen(object lispObject) (lispObject, error) {
	return es.stub("vector-or-char-table-p") // Source file: data.c
}

func (es *emacsStubs) vectorp_autogen(object lispObject) (lispObject, error) {
	return es.stub("vectorp") // Source file: data.c
}

func (es *emacsStubs) dbusInitBus_autogen(bus, private lispObject) (lispObject, error) {
	return es.stub("dbus--init-bus") // Source file: dbusbind.c
}

func (es *emacsStubs) dbusGetUniqueName_autogen(bus lispObject) (lispObject, error) {
	return es.stub("dbus-get-unique-name") // Source file: dbusbind.c
}

func (es *emacsStubs) dbusMessageInternal_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("dbus-message-internal") // Source file: dbusbind.c
}

func (es *emacsStubs) zlibAvailableP_autogen() (lispObject, error) {
	return es.stub("zlib-available-p") // Source file: decompress.c
}

func (es *emacsStubs) zlibDecompressRegion_autogen(start, end, allow_partial lispObject) (lispObject, error) {
	return es.stub("zlib-decompress-region") // Source file: decompress.c
}

func (es *emacsStubs) directoryFiles_autogen(directory, full, match, nosort, count lispObject) (lispObject, error) {
	return es.stub("directory-files") // Source file: dired.c
}

func (es *emacsStubs) directoryFilesAndAttributes_autogen(directory, full, match, nosort, id_format, count lispObject) (lispObject, error) {
	return es.stub("directory-files-and-attributes") // Source file: dired.c
}

func (es *emacsStubs) fileAttributes_autogen(filename, id_format lispObject) (lispObject, error) {
	return es.stub("file-attributes") // Source file: dired.c
}

func (es *emacsStubs) fileAttributesLessp_autogen(f1, f2 lispObject) (lispObject, error) {
	return es.stub("file-attributes-lessp") // Source file: dired.c
}

func (es *emacsStubs) fileNameAllCompletions_autogen(file, directory lispObject) (lispObject, error) {
	return es.stub("file-name-all-completions") // Source file: dired.c
}

func (es *emacsStubs) fileNameCompletion_autogen(file, directory, predicate lispObject) (lispObject, error) {
	return es.stub("file-name-completion") // Source file: dired.c
}

func (es *emacsStubs) systemGroups_autogen() (lispObject, error) {
	return es.stub("system-groups") // Source file: dired.c
}

func (es *emacsStubs) systemUsers_autogen() (lispObject, error) {
	return es.stub("system-users") // Source file: dired.c
}

func (es *emacsStubs) ding_autogen(arg lispObject) (lispObject, error) {
	return es.stub("ding") // Source file: dispnew.c
}

func (es *emacsStubs) displayUpdateForMouseMovement_autogen(mouse_x, mouse_y lispObject) (lispObject, error) {
	return es.stub("display--update-for-mouse-movement") // Source file: dispnew.c
}

func (es *emacsStubs) dumpRedisplayHistory_autogen() (lispObject, error) {
	return es.stub("dump-redisplay-history") // Source file: dispnew.c
}

func (es *emacsStubs) frameOrBufferChangedP_autogen(variable lispObject) (lispObject, error) {
	return es.stub("frame-or-buffer-changed-p") // Source file: dispnew.c
}

func (es *emacsStubs) internalShowCursor_autogen(window, show lispObject) (lispObject, error) {
	return es.stub("internal-show-cursor") // Source file: dispnew.c
}

func (es *emacsStubs) internalShowCursorP_autogen(window lispObject) (lispObject, error) {
	return es.stub("internal-show-cursor-p") // Source file: dispnew.c
}

func (es *emacsStubs) openTermscript_autogen(file lispObject) (lispObject, error) {
	return es.stub("open-termscript") // Source file: dispnew.c
}

func (es *emacsStubs) redisplay_autogen(force lispObject) (lispObject, error) {
	return es.stub("redisplay") // Source file: dispnew.c
}

func (es *emacsStubs) redrawDisplay_autogen() (lispObject, error) {
	return es.stub("redraw-display") // Source file: dispnew.c
}

func (es *emacsStubs) redrawFrame_autogen(frame lispObject) (lispObject, error) {
	return es.stub("redraw-frame") // Source file: dispnew.c
}

func (es *emacsStubs) sendStringToTerminal_autogen(string, terminal lispObject) (lispObject, error) {
	return es.stub("send-string-to-terminal") // Source file: dispnew.c
}

func (es *emacsStubs) sleepFor_autogen(seconds, milliseconds lispObject) (lispObject, error) {
	return es.stub("sleep-for") // Source file: dispnew.c
}

func (es *emacsStubs) snarfDocumentation_autogen(filename lispObject) (lispObject, error) {
	return es.stub("Snarf-documentation") // Source file: doc.c
}

func (es *emacsStubs) documentation_autogen(function, raw lispObject) (lispObject, error) {
	return es.stub("documentation") // Source file: doc.c
}

func (es *emacsStubs) documentationProperty_autogen(symbol, prop, raw lispObject) (lispObject, error) {
	return es.stub("documentation-property") // Source file: doc.c
}

func (es *emacsStubs) subrDocumentation_autogen(function lispObject) (lispObject, error) {
	return es.stub("internal-subr-documentation") // Source file: doc.c
}

func (es *emacsStubs) textQuotingStyle_autogen() (lispObject, error) {
	return es.stub("text-quoting-style") // Source file: doc.c
}

func (es *emacsStubs) bobp_autogen() (lispObject, error) {
	return es.stub("bobp") // Source file: editfns.c
}

func (es *emacsStubs) bolp_autogen() (lispObject, error) {
	return es.stub("bolp") // Source file: editfns.c
}

func (es *emacsStubs) bufferSize_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("buffer-size") // Source file: editfns.c
}

func (es *emacsStubs) bufferString_autogen() (lispObject, error) {
	return es.stub("buffer-string") // Source file: editfns.c
}

func (es *emacsStubs) bufferSubstring_autogen(start, end lispObject) (lispObject, error) {
	return es.stub("buffer-substring") // Source file: editfns.c
}

func (es *emacsStubs) bufferSubstringNoProperties_autogen(start, end lispObject) (lispObject, error) {
	return es.stub("buffer-substring-no-properties") // Source file: editfns.c
}

func (es *emacsStubs) byteToPosition_autogen(bytepos lispObject) (lispObject, error) {
	return es.stub("byte-to-position") // Source file: editfns.c
}

func (es *emacsStubs) byteToString_autogen(byte lispObject) (lispObject, error) {
	return es.stub("byte-to-string") // Source file: editfns.c
}

func (es *emacsStubs) charAfter_autogen(pos lispObject) (lispObject, error) {
	return es.stub("char-after") // Source file: editfns.c
}

func (es *emacsStubs) charBefore_autogen(pos lispObject) (lispObject, error) {
	return es.stub("char-before") // Source file: editfns.c
}

func (es *emacsStubs) charEqual_autogen(c1, c2 lispObject) (lispObject, error) {
	return es.stub("char-equal") // Source file: editfns.c
}

func (es *emacsStubs) charToString_autogen(character lispObject) (lispObject, error) {
	return es.stub("char-to-string") // Source file: editfns.c
}

func (es *emacsStubs) compareBufferSubstrings_autogen(buffer1, start1, end1, buffer2, start2, end2 lispObject) (lispObject, error) {
	return es.stub("compare-buffer-substrings") // Source file: editfns.c
}

func (es *emacsStubs) constrainToField_autogen(new_pos, old_pos, escape_from_edge, only_in_line, inhibit_capture_property lispObject) (lispObject, error) {
	return es.stub("constrain-to-field") // Source file: editfns.c
}

func (es *emacsStubs) currentMessage_autogen() (lispObject, error) {
	return es.stub("current-message") // Source file: editfns.c
}

func (es *emacsStubs) deleteAndExtractRegion_autogen(start, end lispObject) (lispObject, error) {
	return es.stub("delete-and-extract-region") // Source file: editfns.c
}

func (es *emacsStubs) deleteField_autogen(pos lispObject) (lispObject, error) {
	return es.stub("delete-field") // Source file: editfns.c
}

func (es *emacsStubs) deleteRegion_autogen(start, end lispObject) (lispObject, error) {
	return es.stub("delete-region") // Source file: editfns.c
}

func (es *emacsStubs) emacsPid_autogen() (lispObject, error) {
	return es.stub("emacs-pid") // Source file: editfns.c
}

func (es *emacsStubs) eobp_autogen() (lispObject, error) {
	return es.stub("eobp") // Source file: editfns.c
}

func (es *emacsStubs) eolp_autogen() (lispObject, error) {
	return es.stub("eolp") // Source file: editfns.c
}

func (es *emacsStubs) fieldBeginning_autogen(pos, escape_from_edge, limit lispObject) (lispObject, error) {
	return es.stub("field-beginning") // Source file: editfns.c
}

func (es *emacsStubs) fieldEnd_autogen(pos, escape_from_edge, limit lispObject) (lispObject, error) {
	return es.stub("field-end") // Source file: editfns.c
}

func (es *emacsStubs) fieldString_autogen(pos lispObject) (lispObject, error) {
	return es.stub("field-string") // Source file: editfns.c
}

func (es *emacsStubs) fieldStringNoProperties_autogen(pos lispObject) (lispObject, error) {
	return es.stub("field-string-no-properties") // Source file: editfns.c
}

func (es *emacsStubs) followingChar_autogen() (lispObject, error) {
	return es.stub("following-char") // Source file: editfns.c
}

func (es *emacsStubs) format_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("format") // Source file: editfns.c
}

func (es *emacsStubs) formatMessage_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("format-message") // Source file: editfns.c
}

func (es *emacsStubs) gapPosition_autogen() (lispObject, error) {
	return es.stub("gap-position") // Source file: editfns.c
}

func (es *emacsStubs) gapSize_autogen() (lispObject, error) {
	return es.stub("gap-size") // Source file: editfns.c
}

func (es *emacsStubs) getPosProperty_autogen(position, prop, object lispObject) (lispObject, error) {
	return es.stub("get-pos-property") // Source file: editfns.c
}

func (es *emacsStubs) gotoChar_autogen(position lispObject) (lispObject, error) {
	return es.stub("goto-char") // Source file: editfns.c
}

func (es *emacsStubs) groupGid_autogen() (lispObject, error) {
	return es.stub("group-gid") // Source file: editfns.c
}

func (es *emacsStubs) groupName_autogen(gid lispObject) (lispObject, error) {
	return es.stub("group-name") // Source file: editfns.c
}

func (es *emacsStubs) groupRealGid_autogen() (lispObject, error) {
	return es.stub("group-real-gid") // Source file: editfns.c
}

func (es *emacsStubs) insert_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("insert") // Source file: editfns.c
}

func (es *emacsStubs) insertAndInherit_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("insert-and-inherit") // Source file: editfns.c
}

func (es *emacsStubs) insertBeforeMarkers_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("insert-before-markers") // Source file: editfns.c
}

func (es *emacsStubs) insertAndInheritBeforeMarkers_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("insert-before-markers-and-inherit") // Source file: editfns.c
}

func (es *emacsStubs) insertBufferSubstring_autogen(buffer, start, end lispObject) (lispObject, error) {
	return es.stub("insert-buffer-substring") // Source file: editfns.c
}

func (es *emacsStubs) insertByte_autogen(byte, count, inherit lispObject) (lispObject, error) {
	return es.stub("insert-byte") // Source file: editfns.c
}

func (es *emacsStubs) insertChar_autogen(character, count, inherit lispObject) (lispObject, error) {
	return es.stub("insert-char") // Source file: editfns.c
}

func (es *emacsStubs) internalLabeledNarrowToRegion_autogen(start, end, label lispObject) (lispObject, error) {
	return es.stub("internal--labeled-narrow-to-region") // Source file: editfns.c
}

func (es *emacsStubs) internalLabeledWiden_autogen(label lispObject) (lispObject, error) {
	return es.stub("internal--labeled-widen") // Source file: editfns.c
}

func (es *emacsStubs) lineBeginningPosition_autogen(n lispObject) (lispObject, error) {
	return es.stub("line-beginning-position") // Source file: editfns.c
}

func (es *emacsStubs) lineEndPosition_autogen(n lispObject) (lispObject, error) {
	return es.stub("line-end-position") // Source file: editfns.c
}

func (es *emacsStubs) markMarker_autogen() (lispObject, error) {
	return es.stub("mark-marker") // Source file: editfns.c
}

func (es *emacsStubs) message_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("message") // Source file: editfns.c
}

func (es *emacsStubs) messageBox_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("message-box") // Source file: editfns.c
}

func (es *emacsStubs) messageOrBox_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("message-or-box") // Source file: editfns.c
}

func (es *emacsStubs) narrowToRegion_autogen(start, end lispObject) (lispObject, error) {
	return es.stub("narrow-to-region") // Source file: editfns.c
}

func (es *emacsStubs) ngettext_autogen(msgid, msgid_plural, n lispObject) (lispObject, error) {
	return es.stub("ngettext") // Source file: editfns.c
}

func (es *emacsStubs) point_autogen() (lispObject, error) {
	return es.stub("point") // Source file: editfns.c
}

func (es *emacsStubs) pointMarker_autogen() (lispObject, error) {
	return es.stub("point-marker") // Source file: editfns.c
}

func (es *emacsStubs) pointMax_autogen() (lispObject, error) {
	return es.stub("point-max") // Source file: editfns.c
}

func (es *emacsStubs) pointMaxMarker_autogen() (lispObject, error) {
	return es.stub("point-max-marker") // Source file: editfns.c
}

func (es *emacsStubs) pointMin_autogen() (lispObject, error) {
	return es.stub("point-min") // Source file: editfns.c
}

func (es *emacsStubs) pointMinMarker_autogen() (lispObject, error) {
	return es.stub("point-min-marker") // Source file: editfns.c
}

func (es *emacsStubs) posBol_autogen(n lispObject) (lispObject, error) {
	return es.stub("pos-bol") // Source file: editfns.c
}

func (es *emacsStubs) posEol_autogen(n lispObject) (lispObject, error) {
	return es.stub("pos-eol") // Source file: editfns.c
}

func (es *emacsStubs) positionBytes_autogen(position lispObject) (lispObject, error) {
	return es.stub("position-bytes") // Source file: editfns.c
}

func (es *emacsStubs) previousChar_autogen() (lispObject, error) {
	return es.stub("preceding-char") // Source file: editfns.c
}

func (es *emacsStubs) propertize_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("propertize") // Source file: editfns.c
}

func (es *emacsStubs) regionBeginning_autogen() (lispObject, error) {
	return es.stub("region-beginning") // Source file: editfns.c
}

func (es *emacsStubs) regionEnd_autogen() (lispObject, error) {
	return es.stub("region-end") // Source file: editfns.c
}

func (es *emacsStubs) replaceBufferContents_autogen(source, max_secs, max_costs lispObject) (lispObject, error) {
	return es.stub("replace-buffer-contents") // Source file: editfns.c
}

func (es *emacsStubs) saveCurrentBuffer_autogen(args lispObject) (lispObject, error) {
	return es.stub("save-current-buffer") // Source file: editfns.c
}

func (es *emacsStubs) saveExcursion_autogen(args lispObject) (lispObject, error) {
	return es.stub("save-excursion") // Source file: editfns.c
}

func (es *emacsStubs) saveRestriction_autogen(body lispObject) (lispObject, error) {
	return es.stub("save-restriction") // Source file: editfns.c
}

func (es *emacsStubs) stringToChar_autogen(string lispObject) (lispObject, error) {
	return es.stub("string-to-char") // Source file: editfns.c
}

func (es *emacsStubs) substCharInRegion_autogen(start, end, fromchar, tochar, noundo lispObject) (lispObject, error) {
	return es.stub("subst-char-in-region") // Source file: editfns.c
}

func (es *emacsStubs) systemName_autogen() (lispObject, error) {
	return es.stub("system-name") // Source file: editfns.c
}

func (es *emacsStubs) translateRegionInternal_autogen(start, end, table lispObject) (lispObject, error) {
	return es.stub("translate-region-internal") // Source file: editfns.c
}

func (es *emacsStubs) transposeRegions_autogen(startr1, endr1, startr2, endr2, leave_markers lispObject) (lispObject, error) {
	return es.stub("transpose-regions") // Source file: editfns.c
}

func (es *emacsStubs) userFullName_autogen(uid lispObject) (lispObject, error) {
	return es.stub("user-full-name") // Source file: editfns.c
}

func (es *emacsStubs) userLoginName_autogen(uid lispObject) (lispObject, error) {
	return es.stub("user-login-name") // Source file: editfns.c
}

func (es *emacsStubs) userRealLoginName_autogen() (lispObject, error) {
	return es.stub("user-real-login-name") // Source file: editfns.c
}

func (es *emacsStubs) userRealUid_autogen() (lispObject, error) {
	return es.stub("user-real-uid") // Source file: editfns.c
}

func (es *emacsStubs) userUid_autogen() (lispObject, error) {
	return es.stub("user-uid") // Source file: editfns.c
}

func (es *emacsStubs) widen_autogen() (lispObject, error) {
	return es.stub("widen") // Source file: editfns.c
}

func (es *emacsStubs) moduleLoad_autogen(file lispObject) (lispObject, error) {
	return es.stub("module-load") // Source file: emacs-module.c
}

func (es *emacsStubs) daemonInitialized_autogen() (lispObject, error) {
	return es.stub("daemon-initialized") // Source file: emacs.c
}

func (es *emacsStubs) daemonp_autogen() (lispObject, error) {
	return es.stub("daemonp") // Source file: emacs.c
}

func (es *emacsStubs) dumpEmacs_autogen(filename, symfile lispObject) (lispObject, error) {
	return es.stub("dump-emacs") // Source file: emacs.c
}

func (es *emacsStubs) invocationDirectory_autogen() (lispObject, error) {
	return es.stub("invocation-directory") // Source file: emacs.c
}

func (es *emacsStubs) invocationName_autogen() (lispObject, error) {
	return es.stub("invocation-name") // Source file: emacs.c
}

func (es *emacsStubs) killEmacs_autogen(arg, restart lispObject) (lispObject, error) {
	return es.stub("kill-emacs") // Source file: emacs.c, attributes: noreturn
}

func (es *emacsStubs) and_autogen(args lispObject) (lispObject, error) {
	return es.stub("and") // Source file: eval.c
}

func (es *emacsStubs) apply_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("apply") // Source file: eval.c
}

func (es *emacsStubs) autoload_autogen(function, file, docstring, interactive, type_ lispObject) (lispObject, error) {
	return es.stub("autoload") // Source file: eval.c
}

func (es *emacsStubs) autoloadDoLoad_autogen(fundef, funname, macro_only lispObject) (lispObject, error) {
	return es.stub("autoload-do-load") // Source file: eval.c
}

func (es *emacsStubs) backtraceFramesFromThread_autogen(thread lispObject) (lispObject, error) {
	return es.stub("backtrace--frames-from-thread") // Source file: eval.c
}

func (es *emacsStubs) backtraceLocals_autogen(nframes, base lispObject) (lispObject, error) {
	return es.stub("backtrace--locals") // Source file: eval.c
}

func (es *emacsStubs) backtraceDebug_autogen(level, flag lispObject) (lispObject, error) {
	return es.stub("backtrace-debug") // Source file: eval.c
}

func (es *emacsStubs) backtraceEval_autogen(exp, nframes, base lispObject) (lispObject, error) {
	return es.stub("backtrace-eval") // Source file: eval.c
}

func (es *emacsStubs) backtraceFrameInternal_autogen(function, nframes, base lispObject) (lispObject, error) {
	return es.stub("backtrace-frame--internal") // Source file: eval.c
}

func (es *emacsStubs) catch_autogen(args lispObject) (lispObject, error) {
	return es.stub("catch") // Source file: eval.c
}

func (es *emacsStubs) commandp_autogen(function, for_call_interactively lispObject) (lispObject, error) {
	return es.stub("commandp") // Source file: eval.c
}

func (es *emacsStubs) cond_autogen(args lispObject) (lispObject, error) {
	return es.stub("cond") // Source file: eval.c
}

func (es *emacsStubs) conditionCase_autogen(args lispObject) (lispObject, error) {
	return es.stub("condition-case") // Source file: eval.c
}

func (es *emacsStubs) defaultToplevelValue_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("default-toplevel-value") // Source file: eval.c
}

func (es *emacsStubs) defconst_autogen(args lispObject) (lispObject, error) {
	return es.stub("defconst") // Source file: eval.c
}

func (es *emacsStubs) defconst1_autogen(sym, initvalue, docstring lispObject) (lispObject, error) {
	return es.stub("defconst-1") // Source file: eval.c
}

func (es *emacsStubs) defvar_autogen(args lispObject) (lispObject, error) {
	return es.stub("defvar") // Source file: eval.c
}

func (es *emacsStubs) defvar1_autogen(sym, initvalue, docstring lispObject) (lispObject, error) {
	return es.stub("defvar-1") // Source file: eval.c
}

func (es *emacsStubs) defvaralias_autogen(new_alias, base_variable, docstring lispObject) (lispObject, error) {
	return es.stub("defvaralias") // Source file: eval.c
}

func (es *emacsStubs) eval_autogen(form, lexical lispObject) (lispObject, error) {
	return es.stub("eval") // Source file: eval.c
}

func (es *emacsStubs) fetchBytecode_autogen(object lispObject) (lispObject, error) {
	return es.stub("fetch-bytecode") // Source file: eval.c
}

func (es *emacsStubs) funcArity_autogen(function lispObject) (lispObject, error) {
	return es.stub("func-arity") // Source file: eval.c
}

func (es *emacsStubs) funcall_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("funcall") // Source file: eval.c
}

func (es *emacsStubs) funcallWithDelayedMessage_autogen(timeout, message, function lispObject) (lispObject, error) {
	return es.stub("funcall-with-delayed-message") // Source file: eval.c
}

func (es *emacsStubs) function_autogen(args lispObject) (lispObject, error) {
	return es.stub("function") // Source file: eval.c
}

func (es *emacsStubs) functionp_autogen(object lispObject) (lispObject, error) {
	return es.stub("functionp") // Source file: eval.c
}

func (es *emacsStubs) if_autogen(args lispObject) (lispObject, error) {
	return es.stub("if") // Source file: eval.c
}

func (es *emacsStubs) internalDefineUninitializedVariable_autogen(symbol, doc lispObject) (lispObject, error) {
	return es.stub("internal--define-uninitialized-variable") // Source file: eval.c
}

func (es *emacsStubs) makeVarNonSpecial_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("internal-make-var-non-special") // Source file: eval.c
}

func (es *emacsStubs) let_autogen(args lispObject) (lispObject, error) {
	return es.stub("let") // Source file: eval.c
}

func (es *emacsStubs) letX_autogen(args lispObject) (lispObject, error) {
	return es.stub("let*") // Source file: eval.c
}

func (es *emacsStubs) macroexpand_autogen(form, environment lispObject) (lispObject, error) {
	return es.stub("macroexpand") // Source file: eval.c
}

func (es *emacsStubs) mapbacktrace_autogen(function, base lispObject) (lispObject, error) {
	return es.stub("mapbacktrace") // Source file: eval.c
}

func (es *emacsStubs) or_autogen(args lispObject) (lispObject, error) {
	return es.stub("or") // Source file: eval.c
}

func (es *emacsStubs) prog1_autogen(args lispObject) (lispObject, error) {
	return es.stub("prog1") // Source file: eval.c
}

func (es *emacsStubs) progn_autogen(body lispObject) (lispObject, error) {
	return es.stub("progn") // Source file: eval.c
}

func (es *emacsStubs) quote_autogen(args lispObject) (lispObject, error) {
	return es.stub("quote") // Source file: eval.c
}

func (es *emacsStubs) runHookWithArgs_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("run-hook-with-args") // Source file: eval.c
}

func (es *emacsStubs) runHookWithArgsUntilFailure_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("run-hook-with-args-until-failure") // Source file: eval.c
}

func (es *emacsStubs) runHookWithArgsUntilSuccess_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("run-hook-with-args-until-success") // Source file: eval.c
}

func (es *emacsStubs) runHookWrapped_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("run-hook-wrapped") // Source file: eval.c
}

func (es *emacsStubs) runHooks_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("run-hooks") // Source file: eval.c
}

func (es *emacsStubs) setDefaultToplevelValue_autogen(symbol, value lispObject) (lispObject, error) {
	return es.stub("set-default-toplevel-value") // Source file: eval.c
}

func (es *emacsStubs) setq_autogen(args lispObject) (lispObject, error) {
	return es.stub("setq") // Source file: eval.c
}

func (es *emacsStubs) signal_autogen(error_symbol, data lispObject) (lispObject, error) {
	return es.stub("signal") // Source file: eval.c, attributes: noreturn
}

func (es *emacsStubs) specialVariableP_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("special-variable-p") // Source file: eval.c
}

func (es *emacsStubs) throw_autogen(tag, value lispObject) (lispObject, error) {
	return es.stub("throw") // Source file: eval.c, attributes: noreturn
}

func (es *emacsStubs) unwindProtect_autogen(args lispObject) (lispObject, error) {
	return es.stub("unwind-protect") // Source file: eval.c
}

func (es *emacsStubs) while_autogen(args lispObject) (lispObject, error) {
	return es.stub("while") // Source file: eval.c
}

func (es *emacsStubs) accessFile_autogen(filename, string lispObject) (lispObject, error) {
	return es.stub("access-file") // Source file: fileio.c
}

func (es *emacsStubs) addNameToFile_autogen(file, newname, ok_if_already_exists lispObject) (lispObject, error) {
	return es.stub("add-name-to-file") // Source file: fileio.c
}

func (es *emacsStubs) carLessThanCar_autogen(a, b lispObject) (lispObject, error) {
	return es.stub("car-less-than-car") // Source file: fileio.c
}

func (es *emacsStubs) clearBufferAutoSaveFailure_autogen() (lispObject, error) {
	return es.stub("clear-buffer-auto-save-failure") // Source file: fileio.c
}

func (es *emacsStubs) copyFile_autogen(file, newname, ok_if_already_exists, keep_time, preserve_uid_gid, preserve_permissions lispObject) (lispObject, error) {
	return es.stub("copy-file") // Source file: fileio.c
}

func (es *emacsStubs) defaultFileModes_autogen() (lispObject, error) {
	return es.stub("default-file-modes") // Source file: fileio.c
}

func (es *emacsStubs) deleteDirectoryInternal_autogen(directory lispObject) (lispObject, error) {
	return es.stub("delete-directory-internal") // Source file: fileio.c
}

func (es *emacsStubs) deleteFileInternal_autogen(filename lispObject) (lispObject, error) {
	return es.stub("delete-file-internal") // Source file: fileio.c
}

func (es *emacsStubs) directoryFileName_autogen(directory lispObject) (lispObject, error) {
	return es.stub("directory-file-name") // Source file: fileio.c
}

func (es *emacsStubs) directoryNameP_autogen(name lispObject) (lispObject, error) {
	return es.stub("directory-name-p") // Source file: fileio.c
}

func (es *emacsStubs) doAutoSave_autogen(no_message, current_only lispObject) (lispObject, error) {
	return es.stub("do-auto-save") // Source file: fileio.c
}

func (es *emacsStubs) expandFileName_autogen(name, default_directory lispObject) (lispObject, error) {
	return es.stub("expand-file-name") // Source file: fileio.c
}

func (es *emacsStubs) fileAccessibleDirectoryP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-accessible-directory-p") // Source file: fileio.c
}

func (es *emacsStubs) fileAcl_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-acl") // Source file: fileio.c
}

func (es *emacsStubs) fileDirectoryP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-directory-p") // Source file: fileio.c
}

func (es *emacsStubs) fileExecutableP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-executable-p") // Source file: fileio.c
}

func (es *emacsStubs) fileExistsP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-exists-p") // Source file: fileio.c
}

func (es *emacsStubs) fileModes_autogen(filename, flag lispObject) (lispObject, error) {
	return es.stub("file-modes") // Source file: fileio.c
}

func (es *emacsStubs) fileNameAbsoluteP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-name-absolute-p") // Source file: fileio.c
}

func (es *emacsStubs) fileNameAsDirectory_autogen(file lispObject) (lispObject, error) {
	return es.stub("file-name-as-directory") // Source file: fileio.c
}

func (es *emacsStubs) fileNameCaseInsensitiveP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-name-case-insensitive-p") // Source file: fileio.c
}

func (es *emacsStubs) fileNameConcat_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("file-name-concat") // Source file: fileio.c
}

func (es *emacsStubs) fileNameDirectory_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-name-directory") // Source file: fileio.c
}

func (es *emacsStubs) fileNameNondirectory_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-name-nondirectory") // Source file: fileio.c
}

func (es *emacsStubs) fileNewerThanFileP_autogen(file1, file2 lispObject) (lispObject, error) {
	return es.stub("file-newer-than-file-p") // Source file: fileio.c
}

func (es *emacsStubs) fileReadableP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-readable-p") // Source file: fileio.c
}

func (es *emacsStubs) fileRegularP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-regular-p") // Source file: fileio.c
}

func (es *emacsStubs) fileSelinuxContext_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-selinux-context") // Source file: fileio.c
}

func (es *emacsStubs) fileSymlinkP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-symlink-p") // Source file: fileio.c
}

func (es *emacsStubs) fileSystemInfo_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-system-info") // Source file: fileio.c
}

func (es *emacsStubs) fileSystemInfo_1_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-system-info") // Source file: w32fns.c
}

func (es *emacsStubs) fileWritableP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-writable-p") // Source file: fileio.c
}

func (es *emacsStubs) findFileNameHandler_autogen(filename, operation lispObject) (lispObject, error) {
	return es.stub("find-file-name-handler") // Source file: fileio.c
}

func (es *emacsStubs) insertFileContents_autogen(filename, visit, beg, end, replace lispObject) (lispObject, error) {
	return es.stub("insert-file-contents") // Source file: fileio.c
}

func (es *emacsStubs) makeDirectoryInternal_autogen(directory lispObject) (lispObject, error) {
	return es.stub("make-directory-internal") // Source file: fileio.c
}

func (es *emacsStubs) makeSymbolicLink_autogen(target, linkname, ok_if_already_exists lispObject) (lispObject, error) {
	return es.stub("make-symbolic-link") // Source file: fileio.c
}

func (es *emacsStubs) makeTempFileInternal_autogen(prefix, dir_flag, suffix, text lispObject) (lispObject, error) {
	return es.stub("make-temp-file-internal") // Source file: fileio.c
}

func (es *emacsStubs) makeTempName_autogen(prefix lispObject) (lispObject, error) {
	return es.stub("make-temp-name") // Source file: fileio.c
}

func (es *emacsStubs) nextReadFileUsesDialogP_autogen() (lispObject, error) {
	return es.stub("next-read-file-uses-dialog-p") // Source file: fileio.c
}

func (es *emacsStubs) recentAutoSaveP_autogen() (lispObject, error) {
	return es.stub("recent-auto-save-p") // Source file: fileio.c
}

func (es *emacsStubs) renameFile_autogen(file, newname, ok_if_already_exists lispObject) (lispObject, error) {
	return es.stub("rename-file") // Source file: fileio.c
}

func (es *emacsStubs) setBinaryMode_autogen(stream, mode lispObject) (lispObject, error) {
	return es.stub("set-binary-mode") // Source file: fileio.c
}

func (es *emacsStubs) setBufferAutoSaved_autogen() (lispObject, error) {
	return es.stub("set-buffer-auto-saved") // Source file: fileio.c
}

func (es *emacsStubs) setDefaultFileModes_autogen(mode lispObject) (lispObject, error) {
	return es.stub("set-default-file-modes") // Source file: fileio.c
}

func (es *emacsStubs) setFileAcl_autogen(filename, acl_string lispObject) (lispObject, error) {
	return es.stub("set-file-acl") // Source file: fileio.c
}

func (es *emacsStubs) setFileModes_autogen(filename, mode, flag lispObject) (lispObject, error) {
	return es.stub("set-file-modes") // Source file: fileio.c
}

func (es *emacsStubs) setFileSelinuxContext_autogen(filename, context lispObject) (lispObject, error) {
	return es.stub("set-file-selinux-context") // Source file: fileio.c
}

func (es *emacsStubs) setFileTimes_autogen(filename, timestamp, flag lispObject) (lispObject, error) {
	return es.stub("set-file-times") // Source file: fileio.c
}

func (es *emacsStubs) setVisitedFileModtime_autogen(time_flag lispObject) (lispObject, error) {
	return es.stub("set-visited-file-modtime") // Source file: fileio.c
}

func (es *emacsStubs) substituteInFileName_autogen(filename lispObject) (lispObject, error) {
	return es.stub("substitute-in-file-name") // Source file: fileio.c
}

func (es *emacsStubs) unhandledFileNameDirectory_autogen(filename lispObject) (lispObject, error) {
	return es.stub("unhandled-file-name-directory") // Source file: fileio.c
}

func (es *emacsStubs) unixSync_autogen() (lispObject, error) {
	return es.stub("unix-sync") // Source file: fileio.c
}

func (es *emacsStubs) verifyVisitedFileModtime_autogen(buf lispObject) (lispObject, error) {
	return es.stub("verify-visited-file-modtime") // Source file: fileio.c
}

func (es *emacsStubs) visitedFileModtime_autogen() (lispObject, error) {
	return es.stub("visited-file-modtime") // Source file: fileio.c
}

func (es *emacsStubs) writeRegion_autogen(start, end, filename, append, visit, lockname, mustbenew lispObject) (lispObject, error) {
	return es.stub("write-region") // Source file: fileio.c
}

func (es *emacsStubs) fileLockedP_autogen(filename lispObject) (lispObject, error) {
	return es.stub("file-locked-p") // Source file: filelock.c
}

func (es *emacsStubs) lockBuffer_autogen(file lispObject) (lispObject, error) {
	return es.stub("lock-buffer") // Source file: filelock.c
}

func (es *emacsStubs) lockFile_autogen(file lispObject) (lispObject, error) {
	return es.stub("lock-file") // Source file: filelock.c
}

func (es *emacsStubs) unlockBuffer_autogen() (lispObject, error) {
	return es.stub("unlock-buffer") // Source file: filelock.c
}

func (es *emacsStubs) unlockFile_autogen(file lispObject) (lispObject, error) {
	return es.stub("unlock-file") // Source file: filelock.c
}

func (es *emacsStubs) abs_autogen(arg lispObject) (lispObject, error) {
	return es.stub("abs") // Source file: floatfns.c
}

func (es *emacsStubs) acos_autogen(arg lispObject) (lispObject, error) {
	return es.stub("acos") // Source file: floatfns.c
}

func (es *emacsStubs) asin_autogen(arg lispObject) (lispObject, error) {
	return es.stub("asin") // Source file: floatfns.c
}

func (es *emacsStubs) atan_autogen(y, x lispObject) (lispObject, error) {
	return es.stub("atan") // Source file: floatfns.c
}

func (es *emacsStubs) ceiling_autogen(arg, divisor lispObject) (lispObject, error) {
	return es.stub("ceiling") // Source file: floatfns.c
}

func (es *emacsStubs) copysign_autogen(x1, x2 lispObject) (lispObject, error) {
	return es.stub("copysign") // Source file: floatfns.c
}

func (es *emacsStubs) cos_autogen(arg lispObject) (lispObject, error) {
	return es.stub("cos") // Source file: floatfns.c
}

func (es *emacsStubs) exp_autogen(arg lispObject) (lispObject, error) {
	return es.stub("exp") // Source file: floatfns.c
}

func (es *emacsStubs) expt_autogen(arg1, arg2 lispObject) (lispObject, error) {
	return es.stub("expt") // Source file: floatfns.c
}

func (es *emacsStubs) fceiling_autogen(arg lispObject) (lispObject, error) {
	return es.stub("fceiling") // Source file: floatfns.c
}

func (es *emacsStubs) ffloor_autogen(arg lispObject) (lispObject, error) {
	return es.stub("ffloor") // Source file: floatfns.c
}

func (es *emacsStubs) float_autogen(arg lispObject) (lispObject, error) {
	return es.stub("float") // Source file: floatfns.c
}

func (es *emacsStubs) floor_autogen(arg, divisor lispObject) (lispObject, error) {
	return es.stub("floor") // Source file: floatfns.c
}

func (es *emacsStubs) frexp_autogen(x lispObject) (lispObject, error) {
	return es.stub("frexp") // Source file: floatfns.c
}

func (es *emacsStubs) fround_autogen(arg lispObject) (lispObject, error) {
	return es.stub("fround") // Source file: floatfns.c
}

func (es *emacsStubs) ftruncate_autogen(arg lispObject) (lispObject, error) {
	return es.stub("ftruncate") // Source file: floatfns.c
}

func (es *emacsStubs) isnan_autogen(x lispObject) (lispObject, error) {
	return es.stub("isnan") // Source file: floatfns.c
}

func (es *emacsStubs) ldexp_autogen(sgnfcand, exponent lispObject) (lispObject, error) {
	return es.stub("ldexp") // Source file: floatfns.c
}

func (es *emacsStubs) log_autogen(arg, base lispObject) (lispObject, error) {
	return es.stub("log") // Source file: floatfns.c
}

func (es *emacsStubs) logb_autogen(arg lispObject) (lispObject, error) {
	return es.stub("logb") // Source file: floatfns.c
}

func (es *emacsStubs) round_autogen(arg, divisor lispObject) (lispObject, error) {
	return es.stub("round") // Source file: floatfns.c
}

func (es *emacsStubs) sin_autogen(arg lispObject) (lispObject, error) {
	return es.stub("sin") // Source file: floatfns.c
}

func (es *emacsStubs) sqrt_autogen(arg lispObject) (lispObject, error) {
	return es.stub("sqrt") // Source file: floatfns.c
}

func (es *emacsStubs) tan_autogen(arg lispObject) (lispObject, error) {
	return es.stub("tan") // Source file: floatfns.c
}

func (es *emacsStubs) truncate_autogen(arg, divisor lispObject) (lispObject, error) {
	return es.stub("truncate") // Source file: floatfns.c
}

func (es *emacsStubs) append_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("append") // Source file: fns.c
}

func (es *emacsStubs) assoc_autogen(key, alist, testfn lispObject) (lispObject, error) {
	return es.stub("assoc") // Source file: fns.c
}

func (es *emacsStubs) assq_autogen(key, alist lispObject) (lispObject, error) {
	return es.stub("assq") // Source file: fns.c
}

func (es *emacsStubs) base64DecodeRegion_autogen(beg, end, base64url, ignore_invalid lispObject) (lispObject, error) {
	return es.stub("base64-decode-region") // Source file: fns.c
}

func (es *emacsStubs) base64DecodeString_autogen(string, base64url, ignore_invalid lispObject) (lispObject, error) {
	return es.stub("base64-decode-string") // Source file: fns.c
}

func (es *emacsStubs) base64EncodeRegion_autogen(beg, end, no_line_break lispObject) (lispObject, error) {
	return es.stub("base64-encode-region") // Source file: fns.c
}

func (es *emacsStubs) base64EncodeString_autogen(string, no_line_break lispObject) (lispObject, error) {
	return es.stub("base64-encode-string") // Source file: fns.c
}

func (es *emacsStubs) base64urlEncodeRegion_autogen(beg, end, no_pad lispObject) (lispObject, error) {
	return es.stub("base64url-encode-region") // Source file: fns.c
}

func (es *emacsStubs) base64urlEncodeString_autogen(string, no_pad lispObject) (lispObject, error) {
	return es.stub("base64url-encode-string") // Source file: fns.c
}

func (es *emacsStubs) bufferHash_autogen(buffer_or_name lispObject) (lispObject, error) {
	return es.stub("buffer-hash") // Source file: fns.c
}

func (es *emacsStubs) bufferLineStatistics_autogen(buffer_or_name lispObject) (lispObject, error) {
	return es.stub("buffer-line-statistics") // Source file: fns.c
}

func (es *emacsStubs) clearString_autogen(string lispObject) (lispObject, error) {
	return es.stub("clear-string") // Source file: fns.c
}

func (es *emacsStubs) clrhash_autogen(table lispObject) (lispObject, error) {
	return es.stub("clrhash") // Source file: fns.c
}

func (es *emacsStubs) compareStrings_autogen(str1, start1, end1, str2, start2, end2, ignore_case lispObject) (lispObject, error) {
	return es.stub("compare-strings") // Source file: fns.c
}

func (es *emacsStubs) concat_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("concat") // Source file: fns.c
}

func (es *emacsStubs) copyAlist_autogen(alist lispObject) (lispObject, error) {
	return es.stub("copy-alist") // Source file: fns.c
}

func (es *emacsStubs) copyHashTable_autogen(table lispObject) (lispObject, error) {
	return es.stub("copy-hash-table") // Source file: fns.c
}

func (es *emacsStubs) copySequence_autogen(arg lispObject) (lispObject, error) {
	return es.stub("copy-sequence") // Source file: fns.c
}

func (es *emacsStubs) defineHashTableTest_autogen(name, test, hash lispObject) (lispObject, error) {
	return es.stub("define-hash-table-test") // Source file: fns.c
}

func (es *emacsStubs) delete_autogen(elt, seq lispObject) (lispObject, error) {
	return es.stub("delete") // Source file: fns.c
}

func (es *emacsStubs) delq_autogen(elt, list lispObject) (lispObject, error) {
	return es.stub("delq") // Source file: fns.c
}

func (es *emacsStubs) elt_autogen(sequence, n lispObject) (lispObject, error) {
	return es.stub("elt") // Source file: fns.c
}

func (es *emacsStubs) eql_autogen(obj1, obj2 lispObject) (lispObject, error) {
	return es.stub("eql") // Source file: fns.c
}

func (es *emacsStubs) equal_autogen(o1, o2 lispObject) (lispObject, error) {
	return es.stub("equal") // Source file: fns.c
}

func (es *emacsStubs) equalIncludingProperties_autogen(o1, o2 lispObject) (lispObject, error) {
	return es.stub("equal-including-properties") // Source file: fns.c
}

func (es *emacsStubs) featurep_autogen(feature, subfeature lispObject) (lispObject, error) {
	return es.stub("featurep") // Source file: fns.c
}

func (es *emacsStubs) fillarray_autogen(array, item lispObject) (lispObject, error) {
	return es.stub("fillarray") // Source file: fns.c
}

func (es *emacsStubs) get_autogen(symbol, propname lispObject) (lispObject, error) {
	return es.stub("get") // Source file: fns.c
}

func (es *emacsStubs) gethash_autogen(key, table, dflt lispObject) (lispObject, error) {
	return es.stub("gethash") // Source file: fns.c
}

func (es *emacsStubs) hashTableCount_autogen(table lispObject) (lispObject, error) {
	return es.stub("hash-table-count") // Source file: fns.c
}

func (es *emacsStubs) hashTableP_autogen(obj lispObject) (lispObject, error) {
	return es.stub("hash-table-p") // Source file: fns.c
}

func (es *emacsStubs) hashTableRehashSize_autogen(table lispObject) (lispObject, error) {
	return es.stub("hash-table-rehash-size") // Source file: fns.c
}

func (es *emacsStubs) hashTableRehashThreshold_autogen(table lispObject) (lispObject, error) {
	return es.stub("hash-table-rehash-threshold") // Source file: fns.c
}

func (es *emacsStubs) hashTableSize_autogen(table lispObject) (lispObject, error) {
	return es.stub("hash-table-size") // Source file: fns.c
}

func (es *emacsStubs) hashTableTest_autogen(table lispObject) (lispObject, error) {
	return es.stub("hash-table-test") // Source file: fns.c
}

func (es *emacsStubs) hashTableWeakness_autogen(table lispObject) (lispObject, error) {
	return es.stub("hash-table-weakness") // Source file: fns.c
}

func (es *emacsStubs) identity_autogen(argument lispObject) (lispObject, error) {
	return es.stub("identity") // Source file: fns.c, attributes: const
}

func (es *emacsStubs) length_autogen(sequence lispObject) (lispObject, error) {
	return es.stub("length") // Source file: fns.c
}

func (es *emacsStubs) lengthLess_autogen(sequence, length lispObject) (lispObject, error) {
	return es.stub("length<") // Source file: fns.c
}

func (es *emacsStubs) lengthEqual_autogen(sequence, length lispObject) (lispObject, error) {
	return es.stub("length=") // Source file: fns.c
}

func (es *emacsStubs) lengthGreater_autogen(sequence, length lispObject) (lispObject, error) {
	return es.stub("length>") // Source file: fns.c
}

func (es *emacsStubs) lineNumberAtPos_autogen(position, absolute lispObject) (lispObject, error) {
	return es.stub("line-number-at-pos") // Source file: fns.c
}

func (es *emacsStubs) loadAverage_autogen(use_floats lispObject) (lispObject, error) {
	return es.stub("load-average") // Source file: fns.c
}

func (es *emacsStubs) localeInfo_autogen(item lispObject) (lispObject, error) {
	return es.stub("locale-info") // Source file: fns.c
}

func (es *emacsStubs) makeHashTable_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("make-hash-table") // Source file: fns.c
}

func (es *emacsStubs) mapc_autogen(function, sequence lispObject) (lispObject, error) {
	return es.stub("mapc") // Source file: fns.c
}

func (es *emacsStubs) mapcan_autogen(function, sequence lispObject) (lispObject, error) {
	return es.stub("mapcan") // Source file: fns.c
}

func (es *emacsStubs) mapcar_autogen(function, sequence lispObject) (lispObject, error) {
	return es.stub("mapcar") // Source file: fns.c
}

func (es *emacsStubs) mapconcat_autogen(function, sequence, separator lispObject) (lispObject, error) {
	return es.stub("mapconcat") // Source file: fns.c
}

func (es *emacsStubs) maphash_autogen(function, table lispObject) (lispObject, error) {
	return es.stub("maphash") // Source file: fns.c
}

func (es *emacsStubs) md5_autogen(object, start, end, coding_system, noerror lispObject) (lispObject, error) {
	return es.stub("md5") // Source file: fns.c
}

func (es *emacsStubs) member_autogen(elt, list lispObject) (lispObject, error) {
	return es.stub("member") // Source file: fns.c
}

func (es *emacsStubs) memq_autogen(elt, list lispObject) (lispObject, error) {
	return es.stub("memq") // Source file: fns.c
}

func (es *emacsStubs) memql_autogen(elt, list lispObject) (lispObject, error) {
	return es.stub("memql") // Source file: fns.c
}

func (es *emacsStubs) nconc_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("nconc") // Source file: fns.c
}

func (es *emacsStubs) nreverse_autogen(seq lispObject) (lispObject, error) {
	return es.stub("nreverse") // Source file: fns.c
}

func (es *emacsStubs) ntake_autogen(n, list lispObject) (lispObject, error) {
	return es.stub("ntake") // Source file: fns.c
}

func (es *emacsStubs) nth_autogen(n, list lispObject) (lispObject, error) {
	return es.stub("nth") // Source file: fns.c
}

func (es *emacsStubs) nthcdr_autogen(n, list lispObject) (lispObject, error) {
	return es.stub("nthcdr") // Source file: fns.c
}

func (es *emacsStubs) objectIntervals_autogen(object lispObject) (lispObject, error) {
	return es.stub("object-intervals") // Source file: fns.c
}

func (es *emacsStubs) plistGet_autogen(plist, prop, predicate lispObject) (lispObject, error) {
	return es.stub("plist-get") // Source file: fns.c
}

func (es *emacsStubs) plistMember_autogen(plist, prop, predicate lispObject) (lispObject, error) {
	return es.stub("plist-member") // Source file: fns.c
}

func (es *emacsStubs) plistPut_autogen(plist, prop, val, predicate lispObject) (lispObject, error) {
	return es.stub("plist-put") // Source file: fns.c
}

func (es *emacsStubs) properListP_autogen(object lispObject) (lispObject, error) {
	return es.stub("proper-list-p") // Source file: fns.c, attributes: const
}

func (es *emacsStubs) provide_autogen(feature, subfeatures lispObject) (lispObject, error) {
	return es.stub("provide") // Source file: fns.c
}

func (es *emacsStubs) put_autogen(symbol, propname, value lispObject) (lispObject, error) {
	return es.stub("put") // Source file: fns.c
}

func (es *emacsStubs) puthash_autogen(key, value, table lispObject) (lispObject, error) {
	return es.stub("puthash") // Source file: fns.c
}

func (es *emacsStubs) random_autogen(limit lispObject) (lispObject, error) {
	return es.stub("random") // Source file: fns.c
}

func (es *emacsStubs) rassoc_autogen(key, alist lispObject) (lispObject, error) {
	return es.stub("rassoc") // Source file: fns.c
}

func (es *emacsStubs) rassq_autogen(key, alist lispObject) (lispObject, error) {
	return es.stub("rassq") // Source file: fns.c
}

func (es *emacsStubs) remhash_autogen(key, table lispObject) (lispObject, error) {
	return es.stub("remhash") // Source file: fns.c
}

func (es *emacsStubs) require_autogen(feature, filename, noerror lispObject) (lispObject, error) {
	return es.stub("require") // Source file: fns.c
}

func (es *emacsStubs) reverse_autogen(seq lispObject) (lispObject, error) {
	return es.stub("reverse") // Source file: fns.c
}

func (es *emacsStubs) safeLength_autogen(list lispObject) (lispObject, error) {
	return es.stub("safe-length") // Source file: fns.c
}

func (es *emacsStubs) secureHash_autogen(algorithm, object, start, end, binary lispObject) (lispObject, error) {
	return es.stub("secure-hash") // Source file: fns.c
}

func (es *emacsStubs) secureHashAlgorithms_autogen() (lispObject, error) {
	return es.stub("secure-hash-algorithms") // Source file: fns.c
}

func (es *emacsStubs) sort_autogen(seq, predicate lispObject) (lispObject, error) {
	return es.stub("sort") // Source file: fns.c
}

func (es *emacsStubs) stringAsMultibyte_autogen(string lispObject) (lispObject, error) {
	return es.stub("string-as-multibyte") // Source file: fns.c
}

func (es *emacsStubs) stringAsUnibyte_autogen(string lispObject) (lispObject, error) {
	return es.stub("string-as-unibyte") // Source file: fns.c
}

func (es *emacsStubs) stringBytes_autogen(string lispObject) (lispObject, error) {
	return es.stub("string-bytes") // Source file: fns.c
}

func (es *emacsStubs) stringCollateEqualp_autogen(s1, s2, locale, ignore_case lispObject) (lispObject, error) {
	return es.stub("string-collate-equalp") // Source file: fns.c
}

func (es *emacsStubs) stringCollateLessp_autogen(s1, s2, locale, ignore_case lispObject) (lispObject, error) {
	return es.stub("string-collate-lessp") // Source file: fns.c
}

func (es *emacsStubs) stringDistance_autogen(string1, string2, bytecompare lispObject) (lispObject, error) {
	return es.stub("string-distance") // Source file: fns.c
}

func (es *emacsStubs) stringEqual_autogen(s1, s2 lispObject) (lispObject, error) {
	return es.stub("string-equal") // Source file: fns.c
}

func (es *emacsStubs) stringLessp_autogen(string1, string2 lispObject) (lispObject, error) {
	return es.stub("string-lessp") // Source file: fns.c
}

func (es *emacsStubs) stringMakeMultibyte_autogen(string lispObject) (lispObject, error) {
	return es.stub("string-make-multibyte") // Source file: fns.c
}

func (es *emacsStubs) stringMakeUnibyte_autogen(string lispObject) (lispObject, error) {
	return es.stub("string-make-unibyte") // Source file: fns.c
}

func (es *emacsStubs) stringSearch_autogen(needle, haystack, start_pos lispObject) (lispObject, error) {
	return es.stub("string-search") // Source file: fns.c
}

func (es *emacsStubs) stringToMultibyte_autogen(string lispObject) (lispObject, error) {
	return es.stub("string-to-multibyte") // Source file: fns.c
}

func (es *emacsStubs) stringToUnibyte_autogen(string lispObject) (lispObject, error) {
	return es.stub("string-to-unibyte") // Source file: fns.c
}

func (es *emacsStubs) stringVersionLessp_autogen(string1, string2 lispObject) (lispObject, error) {
	return es.stub("string-version-lessp") // Source file: fns.c
}

func (es *emacsStubs) substring_autogen(string, from, to lispObject) (lispObject, error) {
	return es.stub("substring") // Source file: fns.c
}

func (es *emacsStubs) substringNoProperties_autogen(string, from, to lispObject) (lispObject, error) {
	return es.stub("substring-no-properties") // Source file: fns.c
}

func (es *emacsStubs) sxhashEq_autogen(obj lispObject) (lispObject, error) {
	return es.stub("sxhash-eq") // Source file: fns.c
}

func (es *emacsStubs) sxhashEql_autogen(obj lispObject) (lispObject, error) {
	return es.stub("sxhash-eql") // Source file: fns.c
}

func (es *emacsStubs) sxhashEqual_autogen(obj lispObject) (lispObject, error) {
	return es.stub("sxhash-equal") // Source file: fns.c
}

func (es *emacsStubs) sxhashEqualIncludingProperties_autogen(obj lispObject) (lispObject, error) {
	return es.stub("sxhash-equal-including-properties") // Source file: fns.c
}

func (es *emacsStubs) take_autogen(n, list lispObject) (lispObject, error) {
	return es.stub("take") // Source file: fns.c
}

func (es *emacsStubs) vconcat_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("vconcat") // Source file: fns.c
}

func (es *emacsStubs) widgetApply_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("widget-apply") // Source file: fns.c
}

func (es *emacsStubs) widgetGet_autogen(widget, property lispObject) (lispObject, error) {
	return es.stub("widget-get") // Source file: fns.c
}

func (es *emacsStubs) widgetPut_autogen(widget, property, value lispObject) (lispObject, error) {
	return es.stub("widget-put") // Source file: fns.c
}

func (es *emacsStubs) yesOrNoP_autogen(prompt lispObject) (lispObject, error) {
	return es.stub("yes-or-no-p") // Source file: fns.c
}

func (es *emacsStubs) clearFontCache_autogen() (lispObject, error) {
	return es.stub("clear-font-cache") // Source file: font.c
}

func (es *emacsStubs) closeFont_autogen(font_object, frame lispObject) (lispObject, error) {
	return es.stub("close-font") // Source file: font.c
}

func (es *emacsStubs) drawString_autogen(font_object, string lispObject) (lispObject, error) {
	return es.stub("draw-string") // Source file: font.c
}

func (es *emacsStubs) findFont_autogen(font_spec, frame lispObject) (lispObject, error) {
	return es.stub("find-font") // Source file: font.c
}

func (es *emacsStubs) fontAt_autogen(position, window, string lispObject) (lispObject, error) {
	return es.stub("font-at") // Source file: font.c
}

func (es *emacsStubs) fontDriveOtf_autogen(otf_features, gstring_in, from, to, gstring_out, index lispObject) (lispObject, error) {
	return es.stub("font-drive-otf") // Source file: font.c
}

func (es *emacsStubs) fontFaceAttributes_autogen(font, frame lispObject) (lispObject, error) {
	return es.stub("font-face-attributes") // Source file: font.c
}

func (es *emacsStubs) fontFamilyList_autogen(frame lispObject) (lispObject, error) {
	return es.stub("font-family-list") // Source file: font.c
}

func (es *emacsStubs) fontGet_autogen(font, key lispObject) (lispObject, error) {
	return es.stub("font-get") // Source file: font.c
}

func (es *emacsStubs) fontGetGlyphs_autogen(font_object, from, to, object lispObject) (lispObject, error) {
	return es.stub("font-get-glyphs") // Source file: font.c
}

func (es *emacsStubs) fontHasCharP_autogen(font, ch, frame lispObject) (lispObject, error) {
	return es.stub("font-has-char-p") // Source file: font.c
}

func (es *emacsStubs) fontInfo_autogen(name, frame lispObject) (lispObject, error) {
	return es.stub("font-info") // Source file: font.c
}

func (es *emacsStubs) fontMatchP_autogen(spec, font lispObject) (lispObject, error) {
	return es.stub("font-match-p") // Source file: font.c
}

func (es *emacsStubs) fontOtfAlternates_autogen(font_object, character, otf_features lispObject) (lispObject, error) {
	return es.stub("font-otf-alternates") // Source file: font.c
}

func (es *emacsStubs) fontPut_autogen(font, prop, val lispObject) (lispObject, error) {
	return es.stub("font-put") // Source file: font.c
}

func (es *emacsStubs) fontShapeGstring_autogen(gstring, direction lispObject) (lispObject, error) {
	return es.stub("font-shape-gstring") // Source file: font.c
}

func (es *emacsStubs) fontSpec_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("font-spec") // Source file: font.c
}

func (es *emacsStubs) fontVariationGlyphs_autogen(font_object, character lispObject) (lispObject, error) {
	return es.stub("font-variation-glyphs") // Source file: font.c
}

func (es *emacsStubs) fontXlfdName_autogen(font, fold_wildcards lispObject) (lispObject, error) {
	return es.stub("font-xlfd-name") // Source file: font.c
}

func (es *emacsStubs) fontp_autogen(object, extra_type lispObject) (lispObject, error) {
	return es.stub("fontp") // Source file: font.c
}

func (es *emacsStubs) frameFontCache_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-font-cache") // Source file: font.c
}

func (es *emacsStubs) internalCharFont_autogen(position, ch lispObject) (lispObject, error) {
	return es.stub("internal-char-font") // Source file: font.c
}

func (es *emacsStubs) listFonts_autogen(font_spec, frame, num, prefer lispObject) (lispObject, error) {
	return es.stub("list-fonts") // Source file: font.c
}

func (es *emacsStubs) openFont_autogen(font_entity, size, frame lispObject) (lispObject, error) {
	return es.stub("open-font") // Source file: font.c
}

func (es *emacsStubs) queryFont_autogen(font_object lispObject) (lispObject, error) {
	return es.stub("query-font") // Source file: font.c
}

func (es *emacsStubs) fontsetFont_autogen(name, ch, all lispObject) (lispObject, error) {
	return es.stub("fontset-font") // Source file: fontset.c
}

func (es *emacsStubs) fontsetInfo_autogen(fontset, frame lispObject) (lispObject, error) {
	return es.stub("fontset-info") // Source file: fontset.c
}

func (es *emacsStubs) fontsetList_autogen() (lispObject, error) {
	return es.stub("fontset-list") // Source file: fontset.c
}

func (es *emacsStubs) fontsetListAll_autogen() (lispObject, error) {
	return es.stub("fontset-list-all") // Source file: fontset.c
}

func (es *emacsStubs) newFontset_autogen(name, fontlist lispObject) (lispObject, error) {
	return es.stub("new-fontset") // Source file: fontset.c
}

func (es *emacsStubs) queryFontset_autogen(pattern, regexpp lispObject) (lispObject, error) {
	return es.stub("query-fontset") // Source file: fontset.c
}

func (es *emacsStubs) setFontsetFont_autogen(fontset, characters, font_spec, frame, add lispObject) (lispObject, error) {
	return es.stub("set-fontset-font") // Source file: fontset.c
}

func (es *emacsStubs) deleteFrame_autogen(frame, force lispObject) (lispObject, error) {
	return es.stub("delete-frame") // Source file: frame.c
}

func (es *emacsStubs) frameSetWasInvisible_autogen(frame, was_invisible lispObject) (lispObject, error) {
	return es.stub("frame--set-was-invisible") // Source file: frame.c
}

func (es *emacsStubs) frameAfterMakeFrame_autogen(frame, made lispObject) (lispObject, error) {
	return es.stub("frame-after-make-frame") // Source file: frame.c
}

func (es *emacsStubs) frameAncestorP_autogen(ancestor, descendant lispObject) (lispObject, error) {
	return es.stub("frame-ancestor-p") // Source file: frame.c
}

func (es *emacsStubs) bottomDividerWidth_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-bottom-divider-width") // Source file: frame.c
}

func (es *emacsStubs) frameCharHeight_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-char-height") // Source file: frame.c
}

func (es *emacsStubs) frameCharWidth_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-char-width") // Source file: frame.c
}

func (es *emacsStubs) frameChildFrameBorderWidth_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-child-frame-border-width") // Source file: frame.c
}

func (es *emacsStubs) frameFocus_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-focus") // Source file: frame.c
}

func (es *emacsStubs) fringeWidth_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-fringe-width") // Source file: frame.c
}

func (es *emacsStubs) frameInternalBorderWidth_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-internal-border-width") // Source file: frame.c
}

func (es *emacsStubs) frameList_autogen() (lispObject, error) {
	return es.stub("frame-list") // Source file: frame.c
}

func (es *emacsStubs) frameLiveP_autogen(object lispObject) (lispObject, error) {
	return es.stub("frame-live-p") // Source file: frame.c
}

func (es *emacsStubs) frameNativeHeight_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-native-height") // Source file: frame.c
}

func (es *emacsStubs) frameNativeWidth_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-native-width") // Source file: frame.c
}

func (es *emacsStubs) frameParameter_autogen(frame, parameter lispObject) (lispObject, error) {
	return es.stub("frame-parameter") // Source file: frame.c
}

func (es *emacsStubs) frameParameters_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-parameters") // Source file: frame.c
}

func (es *emacsStubs) frameParent_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-parent") // Source file: frame.c
}

func (es *emacsStubs) framePointerVisibleP_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-pointer-visible-p") // Source file: frame.c
}

func (es *emacsStubs) framePosition_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-position") // Source file: frame.c
}

func (es *emacsStubs) rightDividerWidth_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-right-divider-width") // Source file: frame.c
}

func (es *emacsStubs) frameScaleFactor_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-scale-factor") // Source file: frame.c
}

func (es *emacsStubs) scrollBarHeight_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-scroll-bar-height") // Source file: frame.c
}

func (es *emacsStubs) scrollBarWidth_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-scroll-bar-width") // Source file: frame.c
}

func (es *emacsStubs) frameTextCols_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-text-cols") // Source file: frame.c
}

func (es *emacsStubs) frameTextHeight_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-text-height") // Source file: frame.c
}

func (es *emacsStubs) frameTextLines_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-text-lines") // Source file: frame.c
}

func (es *emacsStubs) frameTextWidth_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-text-width") // Source file: frame.c
}

func (es *emacsStubs) frameTotalCols_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-total-cols") // Source file: frame.c
}

func (es *emacsStubs) frameTotalLines_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-total-lines") // Source file: frame.c
}

func (es *emacsStubs) frameVisibleP_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-visible-p") // Source file: frame.c
}

func (es *emacsStubs) frameWindowStateChange_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-window-state-change") // Source file: frame.c
}

func (es *emacsStubs) frameWindowsMinSize_autogen(frame, horizontal, ignore, pixelwise lispObject) (lispObject, error) {
	return es.stub("frame-windows-min-size") // Source file: frame.c, attributes: const
}

func (es *emacsStubs) framep_autogen(object lispObject) (lispObject, error) {
	return es.stub("framep") // Source file: frame.c
}

func (es *emacsStubs) handleSwitchFrame_autogen(event lispObject) (lispObject, error) {
	return es.stub("handle-switch-frame") // Source file: frame.c
}

func (es *emacsStubs) iconifyFrame_autogen(frame lispObject) (lispObject, error) {
	return es.stub("iconify-frame") // Source file: frame.c
}

func (es *emacsStubs) lastNonminibufFrame_autogen() (lispObject, error) {
	return es.stub("last-nonminibuffer-frame") // Source file: frame.c
}

func (es *emacsStubs) lowerFrame_autogen(frame lispObject) (lispObject, error) {
	return es.stub("lower-frame") // Source file: frame.c
}

func (es *emacsStubs) makeFrameInvisible_autogen(frame, force lispObject) (lispObject, error) {
	return es.stub("make-frame-invisible") // Source file: frame.c
}

func (es *emacsStubs) makeFrameVisible_autogen(frame lispObject) (lispObject, error) {
	return es.stub("make-frame-visible") // Source file: frame.c
}

func (es *emacsStubs) makeTerminalFrame_autogen(parms lispObject) (lispObject, error) {
	return es.stub("make-terminal-frame") // Source file: frame.c
}

func (es *emacsStubs) modifyFrameParameters_autogen(frame, alist lispObject) (lispObject, error) {
	return es.stub("modify-frame-parameters") // Source file: frame.c
}

func (es *emacsStubs) mousePixelPosition_autogen() (lispObject, error) {
	return es.stub("mouse-pixel-position") // Source file: frame.c
}

func (es *emacsStubs) mousePosition_autogen() (lispObject, error) {
	return es.stub("mouse-position") // Source file: frame.c
}

func (es *emacsStubs) nextFrame_autogen(frame, miniframe lispObject) (lispObject, error) {
	return es.stub("next-frame") // Source file: frame.c
}

func (es *emacsStubs) oldSelectedFrame_autogen() (lispObject, error) {
	return es.stub("old-selected-frame") // Source file: frame.c
}

func (es *emacsStubs) previousFrame_autogen(frame, miniframe lispObject) (lispObject, error) {
	return es.stub("previous-frame") // Source file: frame.c
}

func (es *emacsStubs) raiseFrame_autogen(frame lispObject) (lispObject, error) {
	return es.stub("raise-frame") // Source file: frame.c
}

func (es *emacsStubs) reconsiderFrameFonts_autogen(frame lispObject) (lispObject, error) {
	return es.stub("reconsider-frame-fonts") // Source file: frame.c
}

func (es *emacsStubs) redirectFrameFocus_autogen(frame, focus_frame lispObject) (lispObject, error) {
	return es.stub("redirect-frame-focus") // Source file: frame.c
}

func (es *emacsStubs) selectFrame_autogen(frame, norecord lispObject) (lispObject, error) {
	return es.stub("select-frame") // Source file: frame.c
}

func (es *emacsStubs) selectedFrame_autogen() (lispObject, error) {
	return es.stub("selected-frame") // Source file: frame.c
}

func (es *emacsStubs) setFrameHeight_autogen(frame, height, pretend, pixelwise lispObject) (lispObject, error) {
	return es.stub("set-frame-height") // Source file: frame.c
}

func (es *emacsStubs) setFramePosition_autogen(frame, x, y lispObject) (lispObject, error) {
	return es.stub("set-frame-position") // Source file: frame.c
}

func (es *emacsStubs) setFrameSize_autogen(frame, width, height, pixelwise lispObject) (lispObject, error) {
	return es.stub("set-frame-size") // Source file: frame.c
}

func (es *emacsStubs) setFrameWidth_autogen(frame, width, pretend, pixelwise lispObject) (lispObject, error) {
	return es.stub("set-frame-width") // Source file: frame.c
}

func (es *emacsStubs) setFrameWindowStateChange_autogen(frame, arg lispObject) (lispObject, error) {
	return es.stub("set-frame-window-state-change") // Source file: frame.c
}

func (es *emacsStubs) setMousePixelPosition_autogen(frame, x, y lispObject) (lispObject, error) {
	return es.stub("set-mouse-pixel-position") // Source file: frame.c
}

func (es *emacsStubs) setMousePosition_autogen(frame, x, y lispObject) (lispObject, error) {
	return es.stub("set-mouse-position") // Source file: frame.c
}

func (es *emacsStubs) toolBarPixelWidth_autogen(frame lispObject) (lispObject, error) {
	return es.stub("tool-bar-pixel-width") // Source file: frame.c
}

func (es *emacsStubs) visibleFrameList_autogen() (lispObject, error) {
	return es.stub("visible-frame-list") // Source file: frame.c
}

func (es *emacsStubs) windowSystem_autogen(frame lispObject) (lispObject, error) {
	return es.stub("window-system") // Source file: frame.c
}

func (es *emacsStubs) xFocusFrame_autogen(frame, noactivate lispObject) (lispObject, error) {
	return es.stub("x-focus-frame") // Source file: frame.c
}

func (es *emacsStubs) xGetResource_autogen(attribute, class, component, subclass lispObject) (lispObject, error) {
	return es.stub("x-get-resource") // Source file: frame.c
}

func (es *emacsStubs) xParseGeometry_autogen(string lispObject) (lispObject, error) {
	return es.stub("x-parse-geometry") // Source file: frame.c
}

func (es *emacsStubs) defineFringeBitmap_autogen(bitmap, bits, height, width, align lispObject) (lispObject, error) {
	return es.stub("define-fringe-bitmap") // Source file: fringe.c
}

func (es *emacsStubs) destroyFringeBitmap_autogen(bitmap lispObject) (lispObject, error) {
	return es.stub("destroy-fringe-bitmap") // Source file: fringe.c
}

func (es *emacsStubs) fringeBitmapsAtPos_autogen(pos, window lispObject) (lispObject, error) {
	return es.stub("fringe-bitmaps-at-pos") // Source file: fringe.c
}

func (es *emacsStubs) setFringeBitmapFace_autogen(bitmap, face lispObject) (lispObject, error) {
	return es.stub("set-fringe-bitmap-face") // Source file: fringe.c
}

func (es *emacsStubs) gfileAddWatch_autogen(file, flags, callback lispObject) (lispObject, error) {
	return es.stub("gfile-add-watch") // Source file: gfilenotify.c
}

func (es *emacsStubs) gfileMonitorName_autogen(watch_descriptor lispObject) (lispObject, error) {
	return es.stub("gfile-monitor-name") // Source file: gfilenotify.c
}

func (es *emacsStubs) gfileRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return es.stub("gfile-rm-watch") // Source file: gfilenotify.c
}

func (es *emacsStubs) gfileValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return es.stub("gfile-valid-p") // Source file: gfilenotify.c
}

func (es *emacsStubs) gnutlsAsynchronousParameters_autogen(proc, params lispObject) (lispObject, error) {
	return es.stub("gnutls-asynchronous-parameters") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsAvailableP_autogen() (lispObject, error) {
	return es.stub("gnutls-available-p") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsBoot_autogen(proc, type_, proplist lispObject) (lispObject, error) {
	return es.stub("gnutls-boot") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsBye_autogen(proc, cont lispObject) (lispObject, error) {
	return es.stub("gnutls-bye") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsCiphers_autogen() (lispObject, error) {
	return es.stub("gnutls-ciphers") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsDeinit_autogen(proc lispObject) (lispObject, error) {
	return es.stub("gnutls-deinit") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsDigests_autogen() (lispObject, error) {
	return es.stub("gnutls-digests") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsErrorFatalp_autogen(err lispObject) (lispObject, error) {
	return es.stub("gnutls-error-fatalp") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsErrorString_autogen(err lispObject) (lispObject, error) {
	return es.stub("gnutls-error-string") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsErrorp_autogen(err lispObject) (lispObject, error) {
	return es.stub("gnutls-errorp") // Source file: gnutls.c, attributes: const
}

func (es *emacsStubs) gnutlsFormatCertificate_autogen(cert lispObject) (lispObject, error) {
	return es.stub("gnutls-format-certificate") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsGetInitstage_autogen(proc lispObject) (lispObject, error) {
	return es.stub("gnutls-get-initstage") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsHashDigest_autogen(digest_method, input lispObject) (lispObject, error) {
	return es.stub("gnutls-hash-digest") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsHashMac_autogen(hash_method, key, input lispObject) (lispObject, error) {
	return es.stub("gnutls-hash-mac") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsMacs_autogen() (lispObject, error) {
	return es.stub("gnutls-macs") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsPeerStatus_autogen(proc lispObject) (lispObject, error) {
	return es.stub("gnutls-peer-status") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsPeerStatusWarningDescribe_autogen(status_symbol lispObject) (lispObject, error) {
	return es.stub("gnutls-peer-status-warning-describe") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsSymmetricDecrypt_autogen(cipher, key, iv, input, aead_auth lispObject) (lispObject, error) {
	return es.stub("gnutls-symmetric-decrypt") // Source file: gnutls.c
}

func (es *emacsStubs) gnutlsSymmetricEncrypt_autogen(cipher, key, iv, input, aead_auth lispObject) (lispObject, error) {
	return es.stub("gnutls-symmetric-encrypt") // Source file: gnutls.c
}

func (es *emacsStubs) clearImageCache_autogen(filter, animation_cache lispObject) (lispObject, error) {
	return es.stub("clear-image-cache") // Source file: image.c
}

func (es *emacsStubs) imageCacheSize_autogen() (lispObject, error) {
	return es.stub("image-cache-size") // Source file: image.c
}

func (es *emacsStubs) imageFlush_autogen(spec, frame lispObject) (lispObject, error) {
	return es.stub("image-flush") // Source file: image.c
}

func (es *emacsStubs) imageMaskP_autogen(spec, frame lispObject) (lispObject, error) {
	return es.stub("image-mask-p") // Source file: image.c
}

func (es *emacsStubs) imageMetadata_autogen(spec, frame lispObject) (lispObject, error) {
	return es.stub("image-metadata") // Source file: image.c
}

func (es *emacsStubs) imageSize_autogen(spec, pixels, frame lispObject) (lispObject, error) {
	return es.stub("image-size") // Source file: image.c
}

func (es *emacsStubs) imageTransformsP_autogen(frame lispObject) (lispObject, error) {
	return es.stub("image-transforms-p") // Source file: image.c
}

func (es *emacsStubs) imagemagickTypes_autogen() (lispObject, error) {
	return es.stub("imagemagick-types") // Source file: image.c
}

func (es *emacsStubs) imagep_autogen(spec lispObject) (lispObject, error) {
	return es.stub("imagep") // Source file: image.c
}

func (es *emacsStubs) initImageLibrary_autogen(type_ lispObject) (lispObject, error) {
	return es.stub("init-image-library") // Source file: image.c
}

func (es *emacsStubs) lookupImage_autogen(spec lispObject) (lispObject, error) {
	return es.stub("lookup-image") // Source file: image.c
}

func (es *emacsStubs) computeMotion_autogen(from, frompos, to, topos, width, offsets, window lispObject) (lispObject, error) {
	return es.stub("compute-motion") // Source file: indent.c
}

func (es *emacsStubs) currentColumn_autogen() (lispObject, error) {
	return es.stub("current-column") // Source file: indent.c
}

func (es *emacsStubs) currentIndentation_autogen() (lispObject, error) {
	return es.stub("current-indentation") // Source file: indent.c
}

func (es *emacsStubs) indentTo_autogen(column, minimum lispObject) (lispObject, error) {
	return es.stub("indent-to") // Source file: indent.c
}

func (es *emacsStubs) lineNumberDisplayWidth_autogen(pixelwise lispObject) (lispObject, error) {
	return es.stub("line-number-display-width") // Source file: indent.c
}

func (es *emacsStubs) moveToColumn_autogen(column, force lispObject) (lispObject, error) {
	return es.stub("move-to-column") // Source file: indent.c
}

func (es *emacsStubs) verticalMotion_autogen(lines, window, cur_col lispObject) (lispObject, error) {
	return es.stub("vertical-motion") // Source file: indent.c
}

func (es *emacsStubs) inotifyAddWatch_autogen(filename, aspect, callback lispObject) (lispObject, error) {
	return es.stub("inotify-add-watch") // Source file: inotify.c
}

func (es *emacsStubs) inotifyAllocatedP_autogen() (lispObject, error) {
	return es.stub("inotify-allocated-p") // Source file: inotify.c
}

func (es *emacsStubs) inotifyRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return es.stub("inotify-rm-watch") // Source file: inotify.c
}

func (es *emacsStubs) inotifyValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return es.stub("inotify-valid-p") // Source file: inotify.c
}

func (es *emacsStubs) inotifyWatchList_autogen() (lispObject, error) {
	return es.stub("inotify-watch-list") // Source file: inotify.c
}

func (es *emacsStubs) combineAfterChangeExecute_autogen() (lispObject, error) {
	return es.stub("combine-after-change-execute") // Source file: insdel.c
}

func (es *emacsStubs) jsonAvailableP_autogen() (lispObject, error) {
	return es.stub("json--available-p") // Source file: json.c
}

func (es *emacsStubs) jsonInsert_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("json-insert") // Source file: json.c
}

func (es *emacsStubs) jsonParseBuffer_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("json-parse-buffer") // Source file: json.c
}

func (es *emacsStubs) jsonParseString_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("json-parse-string") // Source file: json.c
}

func (es *emacsStubs) jsonSerialize_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("json-serialize") // Source file: json.c
}

func (es *emacsStubs) abortRecursiveEdit_autogen() (lispObject, error) {
	return es.stub("abort-recursive-edit") // Source file: keyboard.c, attributes: noreturn
}

func (es *emacsStubs) clearThisCommandKeys_autogen(keep_record lispObject) (lispObject, error) {
	return es.stub("clear-this-command-keys") // Source file: keyboard.c
}

func (es *emacsStubs) commandErrorDefaultFunction_autogen(data, context, signal lispObject) (lispObject, error) {
	return es.stub("command-error-default-function") // Source file: keyboard.c
}

func (es *emacsStubs) currentIdleTime_autogen() (lispObject, error) {
	return es.stub("current-idle-time") // Source file: keyboard.c
}

func (es *emacsStubs) currentInputMode_autogen() (lispObject, error) {
	return es.stub("current-input-mode") // Source file: keyboard.c
}

func (es *emacsStubs) discardInput_autogen() (lispObject, error) {
	return es.stub("discard-input") // Source file: keyboard.c
}

func (es *emacsStubs) eventConvertList_autogen(event_desc lispObject) (lispObject, error) {
	return es.stub("event-convert-list") // Source file: keyboard.c
}

func (es *emacsStubs) exitRecursiveEdit_autogen() (lispObject, error) {
	return es.stub("exit-recursive-edit") // Source file: keyboard.c, attributes: noreturn
}

func (es *emacsStubs) inputPendingP_autogen(check_timers lispObject) (lispObject, error) {
	return es.stub("input-pending-p") // Source file: keyboard.c
}

func (es *emacsStubs) internalTrackMouse_autogen(bodyfun lispObject) (lispObject, error) {
	return es.stub("internal--track-mouse") // Source file: keyboard.c
}

func (es *emacsStubs) eventSymbolParseModifiers_autogen(symbol lispObject) (lispObject, error) {
	return es.stub("internal-event-symbol-parse-modifiers") // Source file: keyboard.c
}

func (es *emacsStubs) internalHandleFocusIn_autogen(event lispObject) (lispObject, error) {
	return es.stub("internal-handle-focus-in") // Source file: keyboard.c
}

func (es *emacsStubs) lossageSize_autogen(arg lispObject) (lispObject, error) {
	return es.stub("lossage-size") // Source file: keyboard.c
}

func (es *emacsStubs) openDribbleFile_autogen(file lispObject) (lispObject, error) {
	return es.stub("open-dribble-file") // Source file: keyboard.c
}

func (es *emacsStubs) posnAtPoint_autogen(pos, window lispObject) (lispObject, error) {
	return es.stub("posn-at-point") // Source file: keyboard.c
}

func (es *emacsStubs) posnAtXY_autogen(x, y, frame_or_window, whole lispObject) (lispObject, error) {
	return es.stub("posn-at-x-y") // Source file: keyboard.c
}

func (es *emacsStubs) readKeySequence_autogen(prompt, continue_echo, dont_downcase_last, can_return_switch_frame, cmd_loop, disable_text_conversion lispObject) (lispObject, error) {
	return es.stub("read-key-sequence") // Source file: keyboard.c
}

func (es *emacsStubs) readKeySequenceVector_autogen(prompt, continue_echo, dont_downcase_last, can_return_switch_frame, cmd_loop, disable_text_conversion lispObject) (lispObject, error) {
	return es.stub("read-key-sequence-vector") // Source file: keyboard.c
}

func (es *emacsStubs) recentKeys_autogen(include_cmds lispObject) (lispObject, error) {
	return es.stub("recent-keys") // Source file: keyboard.c
}

func (es *emacsStubs) recursionDepth_autogen() (lispObject, error) {
	return es.stub("recursion-depth") // Source file: keyboard.c
}

func (es *emacsStubs) recursiveEdit_autogen() (lispObject, error) {
	return es.stub("recursive-edit") // Source file: keyboard.c
}

func (es *emacsStubs) setThisCommandKeys_autogen(keys lispObject) (lispObject, error) {
	return es.stub("set--this-command-keys") // Source file: keyboard.c
}

func (es *emacsStubs) setInputInterruptMode_autogen(interrupt lispObject) (lispObject, error) {
	return es.stub("set-input-interrupt-mode") // Source file: keyboard.c
}

func (es *emacsStubs) setInputMetaMode_autogen(meta, terminal lispObject) (lispObject, error) {
	return es.stub("set-input-meta-mode") // Source file: keyboard.c
}

func (es *emacsStubs) setInputMode_autogen(interrupt, flow, meta, quit lispObject) (lispObject, error) {
	return es.stub("set-input-mode") // Source file: keyboard.c
}

func (es *emacsStubs) setOutputFlowControl_autogen(flow, terminal lispObject) (lispObject, error) {
	return es.stub("set-output-flow-control") // Source file: keyboard.c
}

func (es *emacsStubs) setQuitChar_autogen(quit lispObject) (lispObject, error) {
	return es.stub("set-quit-char") // Source file: keyboard.c
}

func (es *emacsStubs) suspendEmacs_autogen(stuffstring lispObject) (lispObject, error) {
	return es.stub("suspend-emacs") // Source file: keyboard.c
}

func (es *emacsStubs) thisCommandKeys_autogen() (lispObject, error) {
	return es.stub("this-command-keys") // Source file: keyboard.c
}

func (es *emacsStubs) thisCommandKeysVector_autogen() (lispObject, error) {
	return es.stub("this-command-keys-vector") // Source file: keyboard.c
}

func (es *emacsStubs) thisSingleCommandKeys_autogen() (lispObject, error) {
	return es.stub("this-single-command-keys") // Source file: keyboard.c
}

func (es *emacsStubs) thisSingleCommandRawKeys_autogen() (lispObject, error) {
	return es.stub("this-single-command-raw-keys") // Source file: keyboard.c
}

func (es *emacsStubs) topLevel_autogen() (lispObject, error) {
	return es.stub("top-level") // Source file: keyboard.c, attributes: noreturn
}

func (es *emacsStubs) accessibleKeymaps_autogen(keymap, prefix lispObject) (lispObject, error) {
	return es.stub("accessible-keymaps") // Source file: keymap.c
}

func (es *emacsStubs) commandRemapping_autogen(command, position, keymaps lispObject) (lispObject, error) {
	return es.stub("command-remapping") // Source file: keymap.c
}

func (es *emacsStubs) copyKeymap_autogen(keymap lispObject) (lispObject, error) {
	return es.stub("copy-keymap") // Source file: keymap.c
}

func (es *emacsStubs) currentActiveMaps_autogen(olp, position lispObject) (lispObject, error) {
	return es.stub("current-active-maps") // Source file: keymap.c
}

func (es *emacsStubs) currentGlobalMap_autogen() (lispObject, error) {
	return es.stub("current-global-map") // Source file: keymap.c
}

func (es *emacsStubs) currentLocalMap_autogen() (lispObject, error) {
	return es.stub("current-local-map") // Source file: keymap.c
}

func (es *emacsStubs) currentMinorModeMaps_autogen() (lispObject, error) {
	return es.stub("current-minor-mode-maps") // Source file: keymap.c
}

func (es *emacsStubs) defineKey_autogen(keymap, key, def, remove lispObject) (lispObject, error) {
	return es.stub("define-key") // Source file: keymap.c
}

func (es *emacsStubs) describeBufferBindings_autogen(buffer, prefix, menus lispObject) (lispObject, error) {
	return es.stub("describe-buffer-bindings") // Source file: keymap.c
}

func (es *emacsStubs) describeVector_autogen(vector, describer lispObject) (lispObject, error) {
	return es.stub("describe-vector") // Source file: keymap.c
}

func (es *emacsStubs) helpDescribeVector_autogen(vector, prefix, describer, partial, shadow, entire_map, mention_shadow lispObject) (lispObject, error) {
	return es.stub("help--describe-vector") // Source file: keymap.c
}

func (es *emacsStubs) keyBinding_autogen(key, accept_default, no_remap, position lispObject) (lispObject, error) {
	return es.stub("key-binding") // Source file: keymap.c
}

func (es *emacsStubs) keyDescription_autogen(keys, prefix lispObject) (lispObject, error) {
	return es.stub("key-description") // Source file: keymap.c
}

func (es *emacsStubs) keymapGetKeyelt_autogen(object, autoload lispObject) (lispObject, error) {
	return es.stub("keymap--get-keyelt") // Source file: keymap.c
}

func (es *emacsStubs) keymapParent_autogen(keymap lispObject) (lispObject, error) {
	return es.stub("keymap-parent") // Source file: keymap.c
}

func (es *emacsStubs) keymapPrompt_autogen(map_ lispObject) (lispObject, error) {
	return es.stub("keymap-prompt") // Source file: keymap.c
}

func (es *emacsStubs) keymapp_autogen(object lispObject) (lispObject, error) {
	return es.stub("keymapp") // Source file: keymap.c
}

func (es *emacsStubs) lookupKey_autogen(keymap, key, accept_default lispObject) (lispObject, error) {
	return es.stub("lookup-key") // Source file: keymap.c
}

func (es *emacsStubs) makeKeymap_autogen(string lispObject) (lispObject, error) {
	return es.stub("make-keymap") // Source file: keymap.c
}

func (es *emacsStubs) makeSparseKeymap_autogen(string lispObject) (lispObject, error) {
	return es.stub("make-sparse-keymap") // Source file: keymap.c
}

func (es *emacsStubs) mapKeymap_autogen(function, keymap, sort_first lispObject) (lispObject, error) {
	return es.stub("map-keymap") // Source file: keymap.c
}

func (es *emacsStubs) mapKeymapInternal_autogen(function, keymap lispObject) (lispObject, error) {
	return es.stub("map-keymap-internal") // Source file: keymap.c
}

func (es *emacsStubs) minorModeKeyBinding_autogen(key, accept_default lispObject) (lispObject, error) {
	return es.stub("minor-mode-key-binding") // Source file: keymap.c
}

func (es *emacsStubs) setKeymapParent_autogen(keymap, parent lispObject) (lispObject, error) {
	return es.stub("set-keymap-parent") // Source file: keymap.c
}

func (es *emacsStubs) singleKeyDescription_autogen(key, no_angles lispObject) (lispObject, error) {
	return es.stub("single-key-description") // Source file: keymap.c
}

func (es *emacsStubs) textCharDescription_autogen(character lispObject) (lispObject, error) {
	return es.stub("text-char-description") // Source file: keymap.c
}

func (es *emacsStubs) useGlobalMap_autogen(keymap lispObject) (lispObject, error) {
	return es.stub("use-global-map") // Source file: keymap.c
}

func (es *emacsStubs) useLocalMap_autogen(keymap lispObject) (lispObject, error) {
	return es.stub("use-local-map") // Source file: keymap.c
}

func (es *emacsStubs) whereIsInternal_autogen(definition, keymap, firstonly, noindirect, no_remap lispObject) (lispObject, error) {
	return es.stub("where-is-internal") // Source file: keymap.c
}

func (es *emacsStubs) kqueueAddWatch_autogen(file, flags, callback lispObject) (lispObject, error) {
	return es.stub("kqueue-add-watch") // Source file: kqueue.c
}

func (es *emacsStubs) kqueueRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return es.stub("kqueue-rm-watch") // Source file: kqueue.c
}

func (es *emacsStubs) kqueueValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return es.stub("kqueue-valid-p") // Source file: kqueue.c
}

func (es *emacsStubs) lcmsCam02Ucs_autogen(color1, color2, whitepoint, view lispObject) (lispObject, error) {
	return es.stub("lcms-cam02-ucs") // Source file: lcms.c
}

func (es *emacsStubs) lcmsCieDe2000_autogen(color1, color2, kL, kC, kH lispObject) (lispObject, error) {
	return es.stub("lcms-cie-de2000") // Source file: lcms.c
}

func (es *emacsStubs) lcmsJabToJch_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return es.stub("lcms-jab->jch") // Source file: lcms.c
}

func (es *emacsStubs) lcmsJchToJab_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return es.stub("lcms-jch->jab") // Source file: lcms.c
}

func (es *emacsStubs) lcmsJchToXyz_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return es.stub("lcms-jch->xyz") // Source file: lcms.c
}

func (es *emacsStubs) lcmsTempToWhitePoint_autogen(temperature lispObject) (lispObject, error) {
	return es.stub("lcms-temp->white-point") // Source file: lcms.c
}

func (es *emacsStubs) lcmsXyzToJch_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return es.stub("lcms-xyz->jch") // Source file: lcms.c
}

func (es *emacsStubs) lcms2AvailableP_autogen() (lispObject, error) {
	return es.stub("lcms2-available-p") // Source file: lcms.c
}

func (es *emacsStubs) evalBuffer_autogen(buffer, printflag, filename, unibyte, do_allow_print lispObject) (lispObject, error) {
	return es.stub("eval-buffer") // Source file: lread.c
}

func (es *emacsStubs) evalRegion_autogen(start, end, printflag, read_function lispObject) (lispObject, error) {
	return es.stub("eval-region") // Source file: lread.c
}

func (es *emacsStubs) getFileChar_autogen() (lispObject, error) {
	return es.stub("get-file-char") // Source file: lread.c
}

func (es *emacsStubs) getLoadSuffixes_autogen() (lispObject, error) {
	return es.stub("get-load-suffixes") // Source file: lread.c
}

func (es *emacsStubs) intern_autogen(string, obarray lispObject) (lispObject, error) {
	return es.stub("intern") // Source file: lread.c
}

func (es *emacsStubs) internSoft_autogen(name, obarray lispObject) (lispObject, error) {
	return es.stub("intern-soft") // Source file: lread.c
}

func (es *emacsStubs) load_autogen(file, noerror, nomessage, nosuffix, must_suffix lispObject) (lispObject, error) {
	return es.stub("load") // Source file: lread.c
}

func (es *emacsStubs) locateFileInternal_autogen(filename, path, suffixes, predicate lispObject) (lispObject, error) {
	return es.stub("locate-file-internal") // Source file: lread.c
}

func (es *emacsStubs) lreadSubstituteObjectInSubtree_autogen(object, placeholder, completed lispObject) (lispObject, error) {
	return es.stub("lread--substitute-object-in-subtree") // Source file: lread.c
}

func (es *emacsStubs) mapatoms_autogen(function, obarray lispObject) (lispObject, error) {
	return es.stub("mapatoms") // Source file: lread.c
}

func (es *emacsStubs) read_autogen(stream lispObject) (lispObject, error) {
	return es.stub("read") // Source file: lread.c
}

func (es *emacsStubs) readChar_autogen(prompt, inherit_input_method, seconds lispObject) (lispObject, error) {
	return es.stub("read-char") // Source file: lread.c
}

func (es *emacsStubs) readCharExclusive_autogen(prompt, inherit_input_method, seconds lispObject) (lispObject, error) {
	return es.stub("read-char-exclusive") // Source file: lread.c
}

func (es *emacsStubs) readEvent_autogen(prompt, inherit_input_method, seconds lispObject) (lispObject, error) {
	return es.stub("read-event") // Source file: lread.c
}

func (es *emacsStubs) readFromString_autogen(string, start, end lispObject) (lispObject, error) {
	return es.stub("read-from-string") // Source file: lread.c
}

func (es *emacsStubs) readPositioningSymbols_autogen(stream lispObject) (lispObject, error) {
	return es.stub("read-positioning-symbols") // Source file: lread.c
}

func (es *emacsStubs) unintern_autogen(name, obarray lispObject) (lispObject, error) {
	return es.stub("unintern") // Source file: lread.c
}

func (es *emacsStubs) callLastKbdMacro_autogen(prefix, loopfunc lispObject) (lispObject, error) {
	return es.stub("call-last-kbd-macro") // Source file: macros.c
}

func (es *emacsStubs) cancelKbdMacroEvents_autogen() (lispObject, error) {
	return es.stub("cancel-kbd-macro-events") // Source file: macros.c
}

func (es *emacsStubs) endKbdMacro_autogen(repeat, loopfunc lispObject) (lispObject, error) {
	return es.stub("end-kbd-macro") // Source file: macros.c
}

func (es *emacsStubs) executeKbdMacro_autogen(macro, count, loopfunc lispObject) (lispObject, error) {
	return es.stub("execute-kbd-macro") // Source file: macros.c
}

func (es *emacsStubs) startKbdMacro_autogen(append, no_exec lispObject) (lispObject, error) {
	return es.stub("start-kbd-macro") // Source file: macros.c
}

func (es *emacsStubs) storeKbdMacroEvent_autogen(event lispObject) (lispObject, error) {
	return es.stub("store-kbd-macro-event") // Source file: macros.c
}

func (es *emacsStubs) copyMarker_autogen(marker, type_ lispObject) (lispObject, error) {
	return es.stub("copy-marker") // Source file: marker.c
}

func (es *emacsStubs) markerBuffer_autogen(marker lispObject) (lispObject, error) {
	return es.stub("marker-buffer") // Source file: marker.c
}

func (es *emacsStubs) markerInsertionType_autogen(marker lispObject) (lispObject, error) {
	return es.stub("marker-insertion-type") // Source file: marker.c
}

func (es *emacsStubs) markerPosition_autogen(marker lispObject) (lispObject, error) {
	return es.stub("marker-position") // Source file: marker.c
}

func (es *emacsStubs) setMarker_autogen(marker, position, buffer lispObject) (lispObject, error) {
	return es.stub("set-marker") // Source file: marker.c
}

func (es *emacsStubs) setMarkerInsertionType_autogen(marker, type_ lispObject) (lispObject, error) {
	return es.stub("set-marker-insertion-type") // Source file: marker.c
}

func (es *emacsStubs) menuBarMenuAtXY_autogen(x, y, frame lispObject) (lispObject, error) {
	return es.stub("menu-bar-menu-at-x-y") // Source file: menu.c
}

func (es *emacsStubs) xPopupDialog_autogen(position, contents, header lispObject) (lispObject, error) {
	return es.stub("x-popup-dialog") // Source file: menu.c
}

func (es *emacsStubs) xPopupMenu_autogen(position, menu lispObject) (lispObject, error) {
	return es.stub("x-popup-menu") // Source file: menu.c
}

func (es *emacsStubs) abortMinibuffers_autogen() (lispObject, error) {
	return es.stub("abort-minibuffers") // Source file: minibuf.c
}

func (es *emacsStubs) activeMinibufferWindow_autogen() (lispObject, error) {
	return es.stub("active-minibuffer-window") // Source file: minibuf.c
}

func (es *emacsStubs) allCompletions_autogen(string, collection, predicate, hide_spaces lispObject) (lispObject, error) {
	return es.stub("all-completions") // Source file: minibuf.c
}

func (es *emacsStubs) assocString_autogen(key, list, case_fold lispObject) (lispObject, error) {
	return es.stub("assoc-string") // Source file: minibuf.c
}

func (es *emacsStubs) completingRead_autogen(prompt, collection, predicate, require_match, initial_input, hist, def, inherit_input_method lispObject) (lispObject, error) {
	return es.stub("completing-read") // Source file: minibuf.c
}

func (es *emacsStubs) innermostMinibufferP_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("innermost-minibuffer-p") // Source file: minibuf.c
}

func (es *emacsStubs) internalCompleteBuffer_autogen(string, predicate, flag lispObject) (lispObject, error) {
	return es.stub("internal-complete-buffer") // Source file: minibuf.c
}

func (es *emacsStubs) minibufferContents_autogen() (lispObject, error) {
	return es.stub("minibuffer-contents") // Source file: minibuf.c
}

func (es *emacsStubs) minibufferContentsNoProperties_autogen() (lispObject, error) {
	return es.stub("minibuffer-contents-no-properties") // Source file: minibuf.c
}

func (es *emacsStubs) minibufferDepth_autogen() (lispObject, error) {
	return es.stub("minibuffer-depth") // Source file: minibuf.c
}

func (es *emacsStubs) minibufferInnermostCommandLoopP_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("minibuffer-innermost-command-loop-p") // Source file: minibuf.c
}

func (es *emacsStubs) minibufferPrompt_autogen() (lispObject, error) {
	return es.stub("minibuffer-prompt") // Source file: minibuf.c
}

func (es *emacsStubs) minibufferPromptEnd_autogen() (lispObject, error) {
	return es.stub("minibuffer-prompt-end") // Source file: minibuf.c
}

func (es *emacsStubs) minibufferp_autogen(buffer, live lispObject) (lispObject, error) {
	return es.stub("minibufferp") // Source file: minibuf.c
}

func (es *emacsStubs) readBuffer_autogen(prompt, def, require_match, predicate lispObject) (lispObject, error) {
	return es.stub("read-buffer") // Source file: minibuf.c
}

func (es *emacsStubs) readCommand_autogen(prompt, default_value lispObject) (lispObject, error) {
	return es.stub("read-command") // Source file: minibuf.c
}

func (es *emacsStubs) readFromMinibuffer_autogen(prompt, initial_contents, keymap, read, hist, default_value, inherit_input_method lispObject) (lispObject, error) {
	return es.stub("read-from-minibuffer") // Source file: minibuf.c
}

func (es *emacsStubs) readFunction_autogen(prompt lispObject) (lispObject, error) {
	return es.stub("read-function") // Source file: minibuf.c
}

func (es *emacsStubs) readString_autogen(prompt, initial_input, history, default_value, inherit_input_method lispObject) (lispObject, error) {
	return es.stub("read-string") // Source file: minibuf.c
}

func (es *emacsStubs) readVariable_autogen(prompt, default_value lispObject) (lispObject, error) {
	return es.stub("read-variable") // Source file: minibuf.c
}

func (es *emacsStubs) setMinibufferWindow_autogen(window lispObject) (lispObject, error) {
	return es.stub("set-minibuffer-window") // Source file: minibuf.c
}

func (es *emacsStubs) testCompletion_autogen(string, collection, predicate lispObject) (lispObject, error) {
	return es.stub("test-completion") // Source file: minibuf.c
}

func (es *emacsStubs) tryCompletion_autogen(string, collection, predicate lispObject) (lispObject, error) {
	return es.stub("try-completion") // Source file: minibuf.c
}

func (es *emacsStubs) dumpEmacsPortable_autogen(filename, track_referrers lispObject) (lispObject, error) {
	return es.stub("dump-emacs-portable") // Source file: pdumper.c
}

func (es *emacsStubs) dumpEmacsPortableSortPredicate_autogen(a, b lispObject) (lispObject, error) {
	return es.stub("dump-emacs-portable--sort-predicate") // Source file: pdumper.c
}

func (es *emacsStubs) dumpEmacsPortableSortPredicateCopied_autogen(a, b lispObject) (lispObject, error) {
	return es.stub("dump-emacs-portable--sort-predicate-copied") // Source file: pdumper.c
}

func (es *emacsStubs) pdumperStats_autogen() (lispObject, error) {
	return es.stub("pdumper-stats") // Source file: pdumper.c
}

func (es *emacsStubs) pgtkBackendDisplayClass_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("pgtk-backend-display-class") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkDisplayMonitorAttributesList_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("pgtk-display-monitor-attributes-list") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkFontName_autogen(name lispObject) (lispObject, error) {
	return es.stub("pgtk-font-name") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkFrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return es.stub("pgtk-frame-edges") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkFrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return es.stub("pgtk-frame-geometry") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkFrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return es.stub("pgtk-frame-restack") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkGetPageSetup_autogen() (lispObject, error) {
	return es.stub("pgtk-get-page-setup") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkMouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return es.stub("pgtk-mouse-absolute-pixel-position") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkPageSetupDialog_autogen() (lispObject, error) {
	return es.stub("pgtk-page-setup-dialog") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkPrintFramesDialog_autogen(frames lispObject) (lispObject, error) {
	return es.stub("pgtk-print-frames-dialog") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkSetMonitorScaleFactor_autogen(monitor_model, scale_factor lispObject) (lispObject, error) {
	return es.stub("pgtk-set-monitor-scale-factor") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkSetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return es.stub("pgtk-set-mouse-absolute-pixel-position") // Source file: pgtkfns.c
}

func (es *emacsStubs) pgtkSetResource_autogen(attribute, value lispObject) (lispObject, error) {
	return es.stub("pgtk-set-resource") // Source file: pgtkfns.c
}

func (es *emacsStubs) xCloseConnection_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-close-connection") // Source file: pgtkfns.c
}

func (es *emacsStubs) xCloseConnection_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-close-connection") // Source file: w32fns.c
}

func (es *emacsStubs) xCloseConnection_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-close-connection") // Source file: xfns.c
}

func (es *emacsStubs) xCreateFrame_autogen(parms lispObject) (lispObject, error) {
	return es.stub("x-create-frame") // Source file: pgtkfns.c
}

func (es *emacsStubs) xCreateFrame_1_autogen(parameters lispObject) (lispObject, error) {
	return es.stub("x-create-frame") // Source file: w32fns.c
}

func (es *emacsStubs) xCreateFrame_2_autogen(parms lispObject) (lispObject, error) {
	return es.stub("x-create-frame") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayBackingStore_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-backing-store") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayBackingStore_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-backing-store") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayBackingStore_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-backing-store") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayColorCells_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-color-cells") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayColorCells_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-color-cells") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayColorCells_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-color-cells") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayGrayscaleP_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-grayscale-p") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayGrayscaleP_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-grayscale-p") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayGrayscaleP_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-grayscale-p") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayList_autogen() (lispObject, error) {
	return es.stub("x-display-list") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayList_1_autogen() (lispObject, error) {
	return es.stub("x-display-list") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayList_2_autogen() (lispObject, error) {
	return es.stub("x-display-list") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayMmHeight_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-mm-height") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayMmHeight_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-mm-height") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayMmHeight_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-mm-height") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayMmWidth_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-mm-width") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayMmWidth_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-mm-width") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayMmWidth_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-mm-width") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayPixelHeight_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-pixel-height") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayPixelHeight_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-pixel-height") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayPixelHeight_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-pixel-height") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayPixelWidth_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-pixel-width") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayPixelWidth_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-pixel-width") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayPixelWidth_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-pixel-width") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayPlanes_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-planes") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayPlanes_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-planes") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayPlanes_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-planes") // Source file: xfns.c
}

func (es *emacsStubs) xDisplaySaveUnder_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-save-under") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplaySaveUnder_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-save-under") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplaySaveUnder_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-save-under") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayScreens_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-screens") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayScreens_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-screens") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayScreens_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-screens") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayVisualClass_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-visual-class") // Source file: pgtkfns.c
}

func (es *emacsStubs) xDisplayVisualClass_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-display-visual-class") // Source file: w32fns.c
}

func (es *emacsStubs) xDisplayVisualClass_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-visual-class") // Source file: xfns.c
}

func (es *emacsStubs) xExportFrames_autogen(frames, type_ lispObject) (lispObject, error) {
	return es.stub("x-export-frames") // Source file: pgtkfns.c
}

func (es *emacsStubs) xExportFrames_1_autogen(frames, type_ lispObject) (lispObject, error) {
	return es.stub("x-export-frames") // Source file: xfns.c
}

func (es *emacsStubs) xFileDialog_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return es.stub("x-file-dialog") // Source file: pgtkfns.c
}

func (es *emacsStubs) xFileDialog_1_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return es.stub("x-file-dialog") // Source file: w32fns.c
}

func (es *emacsStubs) xFileDialog_2_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return es.stub("x-file-dialog") // Source file: xfns.c
}

func (es *emacsStubs) xFileDialog_3_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return es.stub("x-file-dialog") // Source file: xfns.c
}

func (es *emacsStubs) xGtkDebug_autogen(enable lispObject) (lispObject, error) {
	return es.stub("x-gtk-debug") // Source file: pgtkfns.c
}

func (es *emacsStubs) xGtkDebug_1_autogen(enable lispObject) (lispObject, error) {
	return es.stub("x-gtk-debug") // Source file: xfns.c
}

func (es *emacsStubs) xHideTip_autogen() (lispObject, error) {
	return es.stub("x-hide-tip") // Source file: pgtkfns.c
}

func (es *emacsStubs) xHideTip_1_autogen() (lispObject, error) {
	return es.stub("x-hide-tip") // Source file: w32fns.c
}

func (es *emacsStubs) xHideTip_2_autogen() (lispObject, error) {
	return es.stub("x-hide-tip") // Source file: xfns.c
}

func (es *emacsStubs) xOpenConnection_autogen(display, resource_string, must_succeed lispObject) (lispObject, error) {
	return es.stub("x-open-connection") // Source file: pgtkfns.c
}

func (es *emacsStubs) xOpenConnection_1_autogen(display, xrm_string, must_succeed lispObject) (lispObject, error) {
	return es.stub("x-open-connection") // Source file: w32fns.c
}

func (es *emacsStubs) xOpenConnection_2_autogen(display, xrm_string, must_succeed lispObject) (lispObject, error) {
	return es.stub("x-open-connection") // Source file: xfns.c
}

func (es *emacsStubs) xSelectFont_autogen(frame, ignored lispObject) (lispObject, error) {
	return es.stub("x-select-font") // Source file: pgtkfns.c
}

func (es *emacsStubs) xSelectFont_1_autogen(frame, exclude_proportional lispObject) (lispObject, error) {
	return es.stub("x-select-font") // Source file: w32font.c
}

func (es *emacsStubs) xSelectFont_2_autogen(frame, ignored lispObject) (lispObject, error) {
	return es.stub("x-select-font") // Source file: xfns.c
}

func (es *emacsStubs) xServerMaxRequestSize_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-server-max-request-size") // Source file: pgtkfns.c
}

func (es *emacsStubs) xServerMaxRequestSize_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("x-server-max-request-size") // Source file: w32fns.c
}

func (es *emacsStubs) xServerMaxRequestSize_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-server-max-request-size") // Source file: xfns.c
}

func (es *emacsStubs) xShowTip_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return es.stub("x-show-tip") // Source file: pgtkfns.c
}

func (es *emacsStubs) xShowTip_1_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return es.stub("x-show-tip") // Source file: w32fns.c
}

func (es *emacsStubs) xShowTip_2_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return es.stub("x-show-tip") // Source file: xfns.c
}

func (es *emacsStubs) xwColorDefinedP_autogen(color, frame lispObject) (lispObject, error) {
	return es.stub("xw-color-defined-p") // Source file: pgtkfns.c
}

func (es *emacsStubs) xwColorDefinedP_1_autogen(color, frame lispObject) (lispObject, error) {
	return es.stub("xw-color-defined-p") // Source file: w32fns.c
}

func (es *emacsStubs) xwColorDefinedP_2_autogen(color, frame lispObject) (lispObject, error) {
	return es.stub("xw-color-defined-p") // Source file: xfns.c
}

func (es *emacsStubs) xwColorValues_autogen(color, frame lispObject) (lispObject, error) {
	return es.stub("xw-color-values") // Source file: pgtkfns.c
}

func (es *emacsStubs) xwColorValues_1_autogen(color, frame lispObject) (lispObject, error) {
	return es.stub("xw-color-values") // Source file: w32fns.c
}

func (es *emacsStubs) xwColorValues_2_autogen(color, frame lispObject) (lispObject, error) {
	return es.stub("xw-color-values") // Source file: xfns.c
}

func (es *emacsStubs) xwDisplayColorP_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("xw-display-color-p") // Source file: pgtkfns.c
}

func (es *emacsStubs) xwDisplayColorP_1_autogen(display lispObject) (lispObject, error) {
	return es.stub("xw-display-color-p") // Source file: w32fns.c
}

func (es *emacsStubs) xwDisplayColorP_2_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("xw-display-color-p") // Source file: xfns.c
}

func (es *emacsStubs) pgtkUseImContext_autogen(use_p, terminal lispObject) (lispObject, error) {
	return es.stub("pgtk-use-im-context") // Source file: pgtkim.c
}

func (es *emacsStubs) menuOrPopupActiveP_autogen() (lispObject, error) {
	return es.stub("menu-or-popup-active-p") // Source file: pgtkmenu.c
}

func (es *emacsStubs) menuOrPopupActiveP_1_autogen() (lispObject, error) {
	return es.stub("menu-or-popup-active-p") // Source file: w32menu.c
}

func (es *emacsStubs) menuOrPopupActiveP_2_autogen() (lispObject, error) {
	return es.stub("menu-or-popup-active-p") // Source file: xmenu.c
}

func (es *emacsStubs) xMenuBarOpenInternal_autogen(frame lispObject) (lispObject, error) {
	return es.stub("x-menu-bar-open-internal") // Source file: pgtkmenu.c
}

func (es *emacsStubs) xMenuBarOpenInternal_1_autogen(frame lispObject) (lispObject, error) {
	return es.stub("x-menu-bar-open-internal") // Source file: xmenu.c
}

func (es *emacsStubs) xMenuBarOpenInternal_2_autogen(frame lispObject) (lispObject, error) {
	return es.stub("x-menu-bar-open-internal") // Source file: xmenu.c
}

func (es *emacsStubs) pgtkDisownSelectionInternal_autogen(selection, time_object, terminal lispObject) (lispObject, error) {
	return es.stub("pgtk-disown-selection-internal") // Source file: pgtkselect.c
}

func (es *emacsStubs) pgtkDropFinish_autogen(success, timestamp, delete lispObject) (lispObject, error) {
	return es.stub("pgtk-drop-finish") // Source file: pgtkselect.c
}

func (es *emacsStubs) pgtkGetSelectionInternal_autogen(selection_symbol, target_type, time_stamp, terminal lispObject) (lispObject, error) {
	return es.stub("pgtk-get-selection-internal") // Source file: pgtkselect.c
}

func (es *emacsStubs) pgtkOwnSelectionInternal_autogen(selection, value, frame lispObject) (lispObject, error) {
	return es.stub("pgtk-own-selection-internal") // Source file: pgtkselect.c
}

func (es *emacsStubs) pgtkRegisterDndTargets_autogen(frame, targets lispObject) (lispObject, error) {
	return es.stub("pgtk-register-dnd-targets") // Source file: pgtkselect.c
}

func (es *emacsStubs) pgtkSelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return es.stub("pgtk-selection-exists-p") // Source file: pgtkselect.c
}

func (es *emacsStubs) pgtkSelectionOwnerP_autogen(selection, terminal lispObject) (lispObject, error) {
	return es.stub("pgtk-selection-owner-p") // Source file: pgtkselect.c
}

func (es *emacsStubs) pgtkUpdateDropStatus_autogen(action, timestamp lispObject) (lispObject, error) {
	return es.stub("pgtk-update-drop-status") // Source file: pgtkselect.c
}

func (es *emacsStubs) errorMessageString_autogen(obj lispObject) (lispObject, error) {
	return es.stub("error-message-string") // Source file: print.c
}

func (es *emacsStubs) externalDebuggingOutput_autogen(character lispObject) (lispObject, error) {
	return es.stub("external-debugging-output") // Source file: print.c
}

func (es *emacsStubs) flushStandardOutput_autogen() (lispObject, error) {
	return es.stub("flush-standard-output") // Source file: print.c
}

func (es *emacsStubs) prin1_autogen(object, printcharfun, overrides lispObject) (lispObject, error) {
	return es.stub("prin1") // Source file: print.c
}

func (es *emacsStubs) prin1ToString_autogen(object, noescape, overrides lispObject) (lispObject, error) {
	return es.stub("prin1-to-string") // Source file: print.c
}

func (es *emacsStubs) princ_autogen(object, printcharfun lispObject) (lispObject, error) {
	return es.stub("princ") // Source file: print.c
}

func (es *emacsStubs) print_autogen(object, printcharfun lispObject) (lispObject, error) {
	return es.stub("print") // Source file: print.c
}

func (es *emacsStubs) printPreprocess_autogen(object lispObject) (lispObject, error) {
	return es.stub("print--preprocess") // Source file: print.c
}

func (es *emacsStubs) redirectDebuggingOutput_autogen(file, append lispObject) (lispObject, error) {
	return es.stub("redirect-debugging-output") // Source file: print.c
}

func (es *emacsStubs) terpri_autogen(printcharfun, ensure lispObject) (lispObject, error) {
	return es.stub("terpri") // Source file: print.c
}

func (es *emacsStubs) writeChar_autogen(character, printcharfun lispObject) (lispObject, error) {
	return es.stub("write-char") // Source file: print.c
}

func (es *emacsStubs) acceptProcessOutput_autogen(process, seconds, millisec, just_this_one lispObject) (lispObject, error) {
	return es.stub("accept-process-output") // Source file: process.c
}

func (es *emacsStubs) continueProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return es.stub("continue-process") // Source file: process.c
}

func (es *emacsStubs) deleteProcess_autogen(process lispObject) (lispObject, error) {
	return es.stub("delete-process") // Source file: process.c
}

func (es *emacsStubs) formatNetworkAddress_autogen(address, omit_port lispObject) (lispObject, error) {
	return es.stub("format-network-address") // Source file: process.c
}

func (es *emacsStubs) getBufferProcess_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("get-buffer-process") // Source file: process.c
}

func (es *emacsStubs) getProcess_autogen(name lispObject) (lispObject, error) {
	return es.stub("get-process") // Source file: process.c
}

func (es *emacsStubs) internalDefaultInterruptProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return es.stub("internal-default-interrupt-process") // Source file: process.c
}

func (es *emacsStubs) internalDefaultProcessFilter_autogen(proc, text lispObject) (lispObject, error) {
	return es.stub("internal-default-process-filter") // Source file: process.c
}

func (es *emacsStubs) internalDefaultProcessSentinel_autogen(proc, msg lispObject) (lispObject, error) {
	return es.stub("internal-default-process-sentinel") // Source file: process.c
}

func (es *emacsStubs) internalDefaultSignalProcess_autogen(process, sigcode, remote lispObject) (lispObject, error) {
	return es.stub("internal-default-signal-process") // Source file: process.c
}

func (es *emacsStubs) interruptProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return es.stub("interrupt-process") // Source file: process.c
}

func (es *emacsStubs) killProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return es.stub("kill-process") // Source file: process.c
}

func (es *emacsStubs) listSystemProcesses_autogen() (lispObject, error) {
	return es.stub("list-system-processes") // Source file: process.c
}

func (es *emacsStubs) makeNetworkProcess_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("make-network-process") // Source file: process.c
}

func (es *emacsStubs) makePipeProcess_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("make-pipe-process") // Source file: process.c
}

func (es *emacsStubs) makeProcess_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("make-process") // Source file: process.c
}

func (es *emacsStubs) makeSerialProcess_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("make-serial-process") // Source file: process.c
}

func (es *emacsStubs) networkInterfaceInfo_autogen(ifname lispObject) (lispObject, error) {
	return es.stub("network-interface-info") // Source file: process.c
}

func (es *emacsStubs) networkInterfaceList_autogen(full, family lispObject) (lispObject, error) {
	return es.stub("network-interface-list") // Source file: process.c
}

func (es *emacsStubs) networkLookupAddressInfo_autogen(name, family, hint lispObject) (lispObject, error) {
	return es.stub("network-lookup-address-info") // Source file: process.c
}

func (es *emacsStubs) numProcessors_autogen(query lispObject) (lispObject, error) {
	return es.stub("num-processors") // Source file: process.c
}

func (es *emacsStubs) processAttributes_autogen(pid lispObject) (lispObject, error) {
	return es.stub("process-attributes") // Source file: process.c
}

func (es *emacsStubs) processBuffer_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-buffer") // Source file: process.c
}

func (es *emacsStubs) processCodingSystem_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-coding-system") // Source file: process.c
}

func (es *emacsStubs) processCommand_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-command") // Source file: process.c
}

func (es *emacsStubs) processConnection_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-connection") // Source file: process.c
}

func (es *emacsStubs) processContact_autogen(process, key, no_block lispObject) (lispObject, error) {
	return es.stub("process-contact") // Source file: process.c
}

func (es *emacsStubs) processDatagramAddress_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-datagram-address") // Source file: process.c
}

func (es *emacsStubs) processExitStatus_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-exit-status") // Source file: process.c
}

func (es *emacsStubs) processFilter_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-filter") // Source file: process.c
}

func (es *emacsStubs) processId_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-id") // Source file: process.c
}

func (es *emacsStubs) processInheritCodingSystemFlag_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-inherit-coding-system-flag") // Source file: process.c
}

func (es *emacsStubs) processList_autogen() (lispObject, error) {
	return es.stub("process-list") // Source file: process.c
}

func (es *emacsStubs) processMark_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-mark") // Source file: process.c
}

func (es *emacsStubs) processName_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-name") // Source file: process.c
}

func (es *emacsStubs) processPlist_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-plist") // Source file: process.c
}

func (es *emacsStubs) processQueryOnExitFlag_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-query-on-exit-flag") // Source file: process.c
}

func (es *emacsStubs) processRunningChildP_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-running-child-p") // Source file: process.c
}

func (es *emacsStubs) processSendEof_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-send-eof") // Source file: process.c
}

func (es *emacsStubs) processSendRegion_autogen(process, start, end lispObject) (lispObject, error) {
	return es.stub("process-send-region") // Source file: process.c
}

func (es *emacsStubs) processSendString_autogen(process, string lispObject) (lispObject, error) {
	return es.stub("process-send-string") // Source file: process.c
}

func (es *emacsStubs) processSentinel_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-sentinel") // Source file: process.c
}

func (es *emacsStubs) processStatus_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-status") // Source file: process.c
}

func (es *emacsStubs) processThread_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-thread") // Source file: process.c
}

func (es *emacsStubs) processTtyName_autogen(process, stream lispObject) (lispObject, error) {
	return es.stub("process-tty-name") // Source file: process.c
}

func (es *emacsStubs) processType_autogen(process lispObject) (lispObject, error) {
	return es.stub("process-type") // Source file: process.c
}

func (es *emacsStubs) processp_autogen(object lispObject) (lispObject, error) {
	return es.stub("processp") // Source file: process.c
}

func (es *emacsStubs) quitProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return es.stub("quit-process") // Source file: process.c
}

func (es *emacsStubs) serialProcessConfigure_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("serial-process-configure") // Source file: process.c
}

func (es *emacsStubs) setNetworkProcessOption_autogen(process, option, value, no_error lispObject) (lispObject, error) {
	return es.stub("set-network-process-option") // Source file: process.c
}

func (es *emacsStubs) setProcessBuffer_autogen(process, buffer lispObject) (lispObject, error) {
	return es.stub("set-process-buffer") // Source file: process.c
}

func (es *emacsStubs) setProcessCodingSystem_autogen(process, decoding, encoding lispObject) (lispObject, error) {
	return es.stub("set-process-coding-system") // Source file: process.c
}

func (es *emacsStubs) setProcessDatagramAddress_autogen(process, address lispObject) (lispObject, error) {
	return es.stub("set-process-datagram-address") // Source file: process.c
}

func (es *emacsStubs) setProcessFilter_autogen(process, filter lispObject) (lispObject, error) {
	return es.stub("set-process-filter") // Source file: process.c
}

func (es *emacsStubs) setProcessInheritCodingSystemFlag_autogen(process, flag lispObject) (lispObject, error) {
	return es.stub("set-process-inherit-coding-system-flag") // Source file: process.c
}

func (es *emacsStubs) setProcessPlist_autogen(process, plist lispObject) (lispObject, error) {
	return es.stub("set-process-plist") // Source file: process.c
}

func (es *emacsStubs) setProcessQueryOnExitFlag_autogen(process, flag lispObject) (lispObject, error) {
	return es.stub("set-process-query-on-exit-flag") // Source file: process.c
}

func (es *emacsStubs) setProcessSentinel_autogen(process, sentinel lispObject) (lispObject, error) {
	return es.stub("set-process-sentinel") // Source file: process.c
}

func (es *emacsStubs) setProcessThread_autogen(process, thread lispObject) (lispObject, error) {
	return es.stub("set-process-thread") // Source file: process.c
}

func (es *emacsStubs) setProcessWindowSize_autogen(process, height, width lispObject) (lispObject, error) {
	return es.stub("set-process-window-size") // Source file: process.c
}

func (es *emacsStubs) signalNames_autogen() (lispObject, error) {
	return es.stub("signal-names") // Source file: process.c
}

func (es *emacsStubs) signalProcess_autogen(process, sigcode, remote lispObject) (lispObject, error) {
	return es.stub("signal-process") // Source file: process.c
}

func (es *emacsStubs) stopProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return es.stub("stop-process") // Source file: process.c
}

func (es *emacsStubs) waitingForUserInputP_autogen() (lispObject, error) {
	return es.stub("waiting-for-user-input-p") // Source file: process.c
}

func (es *emacsStubs) functionEqual_autogen(f1, f2 lispObject) (lispObject, error) {
	return es.stub("function-equal") // Source file: profiler.c
}

func (es *emacsStubs) profilerCpuLog_autogen() (lispObject, error) {
	return es.stub("profiler-cpu-log") // Source file: profiler.c
}

func (es *emacsStubs) profilerCpuRunningP_autogen() (lispObject, error) {
	return es.stub("profiler-cpu-running-p") // Source file: profiler.c
}

func (es *emacsStubs) profilerCpuStart_autogen(sampling_interval lispObject) (lispObject, error) {
	return es.stub("profiler-cpu-start") // Source file: profiler.c
}

func (es *emacsStubs) profilerCpuStop_autogen() (lispObject, error) {
	return es.stub("profiler-cpu-stop") // Source file: profiler.c
}

func (es *emacsStubs) profilerMemoryLog_autogen() (lispObject, error) {
	return es.stub("profiler-memory-log") // Source file: profiler.c
}

func (es *emacsStubs) profilerMemoryRunningP_autogen() (lispObject, error) {
	return es.stub("profiler-memory-running-p") // Source file: profiler.c
}

func (es *emacsStubs) profilerMemoryStart_autogen() (lispObject, error) {
	return es.stub("profiler-memory-start") // Source file: profiler.c
}

func (es *emacsStubs) profilerMemoryStop_autogen() (lispObject, error) {
	return es.stub("profiler-memory-stop") // Source file: profiler.c
}

func (es *emacsStubs) lookingAt_autogen(regexp, inhibit_modify lispObject) (lispObject, error) {
	return es.stub("looking-at") // Source file: search.c
}

func (es *emacsStubs) matchBeginning_autogen(subexp lispObject) (lispObject, error) {
	return es.stub("match-beginning") // Source file: search.c
}

func (es *emacsStubs) matchData_autogen(integers, reuse, reseat lispObject) (lispObject, error) {
	return es.stub("match-data") // Source file: search.c
}

func (es *emacsStubs) matchDataTranslate_autogen(n lispObject) (lispObject, error) {
	return es.stub("match-data--translate") // Source file: search.c
}

func (es *emacsStubs) matchEnd_autogen(subexp lispObject) (lispObject, error) {
	return es.stub("match-end") // Source file: search.c
}

func (es *emacsStubs) newlineCacheCheck_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("newline-cache-check") // Source file: search.c
}

func (es *emacsStubs) posixLookingAt_autogen(regexp, inhibit_modify lispObject) (lispObject, error) {
	return es.stub("posix-looking-at") // Source file: search.c
}

func (es *emacsStubs) posixSearchBackward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return es.stub("posix-search-backward") // Source file: search.c
}

func (es *emacsStubs) posixSearchForward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return es.stub("posix-search-forward") // Source file: search.c
}

func (es *emacsStubs) posixStringMatch_autogen(regexp, string, start, inhibit_modify lispObject) (lispObject, error) {
	return es.stub("posix-string-match") // Source file: search.c
}

func (es *emacsStubs) reSearchBackward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return es.stub("re-search-backward") // Source file: search.c
}

func (es *emacsStubs) reSearchForward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return es.stub("re-search-forward") // Source file: search.c
}

func (es *emacsStubs) regexpQuote_autogen(string lispObject) (lispObject, error) {
	return es.stub("regexp-quote") // Source file: search.c
}

func (es *emacsStubs) replaceMatch_autogen(newtext, fixedcase, literal, string, subexp lispObject) (lispObject, error) {
	return es.stub("replace-match") // Source file: search.c
}

func (es *emacsStubs) searchBackward_autogen(string, bound, noerror, count lispObject) (lispObject, error) {
	return es.stub("search-backward") // Source file: search.c
}

func (es *emacsStubs) searchForward_autogen(string, bound, noerror, count lispObject) (lispObject, error) {
	return es.stub("search-forward") // Source file: search.c
}

func (es *emacsStubs) setMatchData_autogen(list, reseat lispObject) (lispObject, error) {
	return es.stub("set-match-data") // Source file: search.c
}

func (es *emacsStubs) stringMatch_autogen(regexp, string, start, inhibit_modify lispObject) (lispObject, error) {
	return es.stub("string-match") // Source file: search.c
}

func (es *emacsStubs) playSoundInternal_autogen(sound lispObject) (lispObject, error) {
	return es.stub("play-sound-internal") // Source file: sound.c
}

func (es *emacsStubs) sqliteAvailableP_autogen() (lispObject, error) {
	return es.stub("sqlite-available-p") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteClose_autogen(db lispObject) (lispObject, error) {
	return es.stub("sqlite-close") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteColumns_autogen(set lispObject) (lispObject, error) {
	return es.stub("sqlite-columns") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteCommit_autogen(db lispObject) (lispObject, error) {
	return es.stub("sqlite-commit") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteExecute_autogen(db, query, values lispObject) (lispObject, error) {
	return es.stub("sqlite-execute") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteFinalize_autogen(set lispObject) (lispObject, error) {
	return es.stub("sqlite-finalize") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteLoadExtension_autogen(db, module lispObject) (lispObject, error) {
	return es.stub("sqlite-load-extension") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteMoreP_autogen(set lispObject) (lispObject, error) {
	return es.stub("sqlite-more-p") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteNext_autogen(set lispObject) (lispObject, error) {
	return es.stub("sqlite-next") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteOpen_autogen(file lispObject) (lispObject, error) {
	return es.stub("sqlite-open") // Source file: sqlite.c
}

func (es *emacsStubs) sqlitePragma_autogen(db, pragma lispObject) (lispObject, error) {
	return es.stub("sqlite-pragma") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteRollback_autogen(db lispObject) (lispObject, error) {
	return es.stub("sqlite-rollback") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteSelect_autogen(db, query, values, return_type lispObject) (lispObject, error) {
	return es.stub("sqlite-select") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteTransaction_autogen(db lispObject) (lispObject, error) {
	return es.stub("sqlite-transaction") // Source file: sqlite.c
}

func (es *emacsStubs) sqliteVersion_autogen() (lispObject, error) {
	return es.stub("sqlite-version") // Source file: sqlite.c
}

func (es *emacsStubs) sqlitep_autogen(object lispObject) (lispObject, error) {
	return es.stub("sqlitep") // Source file: sqlite.c
}

func (es *emacsStubs) backwardPrefixChars_autogen() (lispObject, error) {
	return es.stub("backward-prefix-chars") // Source file: syntax.c
}

func (es *emacsStubs) charSyntax_autogen(character lispObject) (lispObject, error) {
	return es.stub("char-syntax") // Source file: syntax.c
}

func (es *emacsStubs) copySyntaxTable_autogen(table lispObject) (lispObject, error) {
	return es.stub("copy-syntax-table") // Source file: syntax.c
}

func (es *emacsStubs) forwardComment_autogen(count lispObject) (lispObject, error) {
	return es.stub("forward-comment") // Source file: syntax.c
}

func (es *emacsStubs) forwardWord_autogen(arg lispObject) (lispObject, error) {
	return es.stub("forward-word") // Source file: syntax.c
}

func (es *emacsStubs) internalDescribeSyntaxValue_autogen(syntax lispObject) (lispObject, error) {
	return es.stub("internal-describe-syntax-value") // Source file: syntax.c
}

func (es *emacsStubs) matchingParen_autogen(character lispObject) (lispObject, error) {
	return es.stub("matching-paren") // Source file: syntax.c
}

func (es *emacsStubs) modifySyntaxEntry_autogen(c, newentry, syntax_table lispObject) (lispObject, error) {
	return es.stub("modify-syntax-entry") // Source file: syntax.c
}

func (es *emacsStubs) parsePartialSexp_autogen(from, to, targetdepth, stopbefore, oldstate, commentstop lispObject) (lispObject, error) {
	return es.stub("parse-partial-sexp") // Source file: syntax.c
}

func (es *emacsStubs) scanLists_autogen(from, count, depth lispObject) (lispObject, error) {
	return es.stub("scan-lists") // Source file: syntax.c
}

func (es *emacsStubs) scanSexps_autogen(from, count lispObject) (lispObject, error) {
	return es.stub("scan-sexps") // Source file: syntax.c
}

func (es *emacsStubs) setSyntaxTable_autogen(table lispObject) (lispObject, error) {
	return es.stub("set-syntax-table") // Source file: syntax.c
}

func (es *emacsStubs) skipCharsBackward_autogen(string, lim lispObject) (lispObject, error) {
	return es.stub("skip-chars-backward") // Source file: syntax.c
}

func (es *emacsStubs) skipCharsForward_autogen(string, lim lispObject) (lispObject, error) {
	return es.stub("skip-chars-forward") // Source file: syntax.c
}

func (es *emacsStubs) skipSyntaxBackward_autogen(syntax, lim lispObject) (lispObject, error) {
	return es.stub("skip-syntax-backward") // Source file: syntax.c
}

func (es *emacsStubs) skipSyntaxForward_autogen(syntax, lim lispObject) (lispObject, error) {
	return es.stub("skip-syntax-forward") // Source file: syntax.c
}

func (es *emacsStubs) standardSyntaxTable_autogen() (lispObject, error) {
	return es.stub("standard-syntax-table") // Source file: syntax.c
}

func (es *emacsStubs) stringToSyntax_autogen(string lispObject) (lispObject, error) {
	return es.stub("string-to-syntax") // Source file: syntax.c
}

func (es *emacsStubs) syntaxClassToChar_autogen(syntax lispObject) (lispObject, error) {
	return es.stub("syntax-class-to-char") // Source file: syntax.c
}

func (es *emacsStubs) syntaxTable_autogen() (lispObject, error) {
	return es.stub("syntax-table") // Source file: syntax.c
}

func (es *emacsStubs) syntaxTableP_autogen(object lispObject) (lispObject, error) {
	return es.stub("syntax-table-p") // Source file: syntax.c
}

func (es *emacsStubs) getInternalRunTime_autogen() (lispObject, error) {
	return es.stub("get-internal-run-time") // Source file: sysdep.c
}

func (es *emacsStubs) controllingTtyP_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("controlling-tty-p") // Source file: term.c
}

func (es *emacsStubs) gpmMouseStart_autogen() (lispObject, error) {
	return es.stub("gpm-mouse-start") // Source file: term.c
}

func (es *emacsStubs) gpmMouseStop_autogen() (lispObject, error) {
	return es.stub("gpm-mouse-stop") // Source file: term.c
}

func (es *emacsStubs) resumeTty_autogen(tty lispObject) (lispObject, error) {
	return es.stub("resume-tty") // Source file: term.c
}

func (es *emacsStubs) suspendTty_autogen(tty lispObject) (lispObject, error) {
	return es.stub("suspend-tty") // Source file: term.c
}

func (es *emacsStubs) ttyOutputBufferSize_autogen(tty lispObject) (lispObject, error) {
	return es.stub("tty--output-buffer-size") // Source file: term.c
}

func (es *emacsStubs) ttySetOutputBufferSize_autogen(size, tty lispObject) (lispObject, error) {
	return es.stub("tty--set-output-buffer-size") // Source file: term.c
}

func (es *emacsStubs) ttyDisplayColorCells_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("tty-display-color-cells") // Source file: term.c
}

func (es *emacsStubs) ttyDisplayColorP_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("tty-display-color-p") // Source file: term.c
}

func (es *emacsStubs) ttyNoUnderline_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("tty-no-underline") // Source file: term.c
}

func (es *emacsStubs) ttyTopFrame_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("tty-top-frame") // Source file: term.c
}

func (es *emacsStubs) ttyType_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("tty-type") // Source file: term.c
}

func (es *emacsStubs) deleteTerminal_autogen(terminal, force lispObject) (lispObject, error) {
	return es.stub("delete-terminal") // Source file: terminal.c
}

func (es *emacsStubs) frameTerminal_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-terminal") // Source file: terminal.c
}

func (es *emacsStubs) setTerminalParameter_autogen(terminal, parameter, value lispObject) (lispObject, error) {
	return es.stub("set-terminal-parameter") // Source file: terminal.c
}

func (es *emacsStubs) terminalList_autogen() (lispObject, error) {
	return es.stub("terminal-list") // Source file: terminal.c
}

func (es *emacsStubs) terminalLiveP_autogen(object lispObject) (lispObject, error) {
	return es.stub("terminal-live-p") // Source file: terminal.c
}

func (es *emacsStubs) terminalName_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("terminal-name") // Source file: terminal.c
}

func (es *emacsStubs) terminalParameter_autogen(terminal, parameter lispObject) (lispObject, error) {
	return es.stub("terminal-parameter") // Source file: terminal.c
}

func (es *emacsStubs) terminalParameters_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("terminal-parameters") // Source file: terminal.c
}

func (es *emacsStubs) setTextConversionStyle_autogen(value, after_key_sequence lispObject) (lispObject, error) {
	return es.stub("set-text-conversion-style") // Source file: textconv.c
}

func (es *emacsStubs) addFaceTextProperty_autogen(start, end, face, append, object lispObject) (lispObject, error) {
	return es.stub("add-face-text-property") // Source file: textprop.c
}

func (es *emacsStubs) addTextProperties_autogen(start, end, properties, object lispObject) (lispObject, error) {
	return es.stub("add-text-properties") // Source file: textprop.c
}

func (es *emacsStubs) getCharProperty_autogen(position, prop, object lispObject) (lispObject, error) {
	return es.stub("get-char-property") // Source file: textprop.c
}

func (es *emacsStubs) getCharPropertyAndOverlay_autogen(position, prop, object lispObject) (lispObject, error) {
	return es.stub("get-char-property-and-overlay") // Source file: textprop.c
}

func (es *emacsStubs) getTextProperty_autogen(position, prop, object lispObject) (lispObject, error) {
	return es.stub("get-text-property") // Source file: textprop.c
}

func (es *emacsStubs) nextCharPropertyChange_autogen(position, limit lispObject) (lispObject, error) {
	return es.stub("next-char-property-change") // Source file: textprop.c
}

func (es *emacsStubs) nextPropertyChange_autogen(position, object, limit lispObject) (lispObject, error) {
	return es.stub("next-property-change") // Source file: textprop.c
}

func (es *emacsStubs) nextSingleCharPropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return es.stub("next-single-char-property-change") // Source file: textprop.c
}

func (es *emacsStubs) nextSinglePropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return es.stub("next-single-property-change") // Source file: textprop.c
}

func (es *emacsStubs) previousCharPropertyChange_autogen(position, limit lispObject) (lispObject, error) {
	return es.stub("previous-char-property-change") // Source file: textprop.c
}

func (es *emacsStubs) previousPropertyChange_autogen(position, object, limit lispObject) (lispObject, error) {
	return es.stub("previous-property-change") // Source file: textprop.c
}

func (es *emacsStubs) previousSingleCharPropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return es.stub("previous-single-char-property-change") // Source file: textprop.c
}

func (es *emacsStubs) previousSinglePropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return es.stub("previous-single-property-change") // Source file: textprop.c
}

func (es *emacsStubs) putTextProperty_autogen(start, end, property, value, object lispObject) (lispObject, error) {
	return es.stub("put-text-property") // Source file: textprop.c
}

func (es *emacsStubs) removeListOfTextProperties_autogen(start, end, list_of_properties, object lispObject) (lispObject, error) {
	return es.stub("remove-list-of-text-properties") // Source file: textprop.c
}

func (es *emacsStubs) removeTextProperties_autogen(start, end, properties, object lispObject) (lispObject, error) {
	return es.stub("remove-text-properties") // Source file: textprop.c
}

func (es *emacsStubs) setTextProperties_autogen(start, end, properties, object lispObject) (lispObject, error) {
	return es.stub("set-text-properties") // Source file: textprop.c
}

func (es *emacsStubs) textPropertiesAt_autogen(position, object lispObject) (lispObject, error) {
	return es.stub("text-properties-at") // Source file: textprop.c
}

func (es *emacsStubs) textPropertyAny_autogen(start, end, property, value, object lispObject) (lispObject, error) {
	return es.stub("text-property-any") // Source file: textprop.c
}

func (es *emacsStubs) textPropertyNotAll_autogen(start, end, property, value, object lispObject) (lispObject, error) {
	return es.stub("text-property-not-all") // Source file: textprop.c
}

func (es *emacsStubs) allThreads_autogen() (lispObject, error) {
	return es.stub("all-threads") // Source file: thread.c
}

func (es *emacsStubs) conditionMutex_autogen(cond lispObject) (lispObject, error) {
	return es.stub("condition-mutex") // Source file: thread.c
}

func (es *emacsStubs) conditionName_autogen(cond lispObject) (lispObject, error) {
	return es.stub("condition-name") // Source file: thread.c
}

func (es *emacsStubs) conditionNotify_autogen(cond, all lispObject) (lispObject, error) {
	return es.stub("condition-notify") // Source file: thread.c
}

func (es *emacsStubs) conditionWait_autogen(cond lispObject) (lispObject, error) {
	return es.stub("condition-wait") // Source file: thread.c
}

func (es *emacsStubs) currentThread_autogen() (lispObject, error) {
	return es.stub("current-thread") // Source file: thread.c
}

func (es *emacsStubs) makeConditionVariable_autogen(mutex, name lispObject) (lispObject, error) {
	return es.stub("make-condition-variable") // Source file: thread.c
}

func (es *emacsStubs) makeMutex_autogen(name lispObject) (lispObject, error) {
	return es.stub("make-mutex") // Source file: thread.c
}

func (es *emacsStubs) makeThread_autogen(function, name lispObject) (lispObject, error) {
	return es.stub("make-thread") // Source file: thread.c
}

func (es *emacsStubs) mutexLock_autogen(mutex lispObject) (lispObject, error) {
	return es.stub("mutex-lock") // Source file: thread.c
}

func (es *emacsStubs) mutexName_autogen(mutex lispObject) (lispObject, error) {
	return es.stub("mutex-name") // Source file: thread.c
}

func (es *emacsStubs) mutexUnlock_autogen(mutex lispObject) (lispObject, error) {
	return es.stub("mutex-unlock") // Source file: thread.c
}

func (es *emacsStubs) threadBlocker_autogen(thread lispObject) (lispObject, error) {
	return es.stub("thread--blocker") // Source file: thread.c
}

func (es *emacsStubs) threadJoin_autogen(thread lispObject) (lispObject, error) {
	return es.stub("thread-join") // Source file: thread.c
}

func (es *emacsStubs) threadLastError_autogen(cleanup lispObject) (lispObject, error) {
	return es.stub("thread-last-error") // Source file: thread.c
}

func (es *emacsStubs) threadLiveP_autogen(thread lispObject) (lispObject, error) {
	return es.stub("thread-live-p") // Source file: thread.c
}

func (es *emacsStubs) threadName_autogen(thread lispObject) (lispObject, error) {
	return es.stub("thread-name") // Source file: thread.c
}

func (es *emacsStubs) threadSignal_autogen(thread, error_symbol, data lispObject) (lispObject, error) {
	return es.stub("thread-signal") // Source file: thread.c
}

func (es *emacsStubs) threadYield_autogen() (lispObject, error) {
	return es.stub("thread-yield") // Source file: thread.c
}

func (es *emacsStubs) currentCpuTime_autogen() (lispObject, error) {
	return es.stub("current-cpu-time") // Source file: timefns.c
}

func (es *emacsStubs) currentTime_autogen() (lispObject, error) {
	return es.stub("current-time") // Source file: timefns.c
}

func (es *emacsStubs) currentTimeString_autogen(specified_time, zone lispObject) (lispObject, error) {
	return es.stub("current-time-string") // Source file: timefns.c
}

func (es *emacsStubs) currentTimeZone_autogen(specified_time, zone lispObject) (lispObject, error) {
	return es.stub("current-time-zone") // Source file: timefns.c
}

func (es *emacsStubs) decodeTime_autogen(specified_time, zone, form lispObject) (lispObject, error) {
	return es.stub("decode-time") // Source file: timefns.c
}

func (es *emacsStubs) encodeTime_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("encode-time") // Source file: timefns.c
}

func (es *emacsStubs) floatTime_autogen(specified_time lispObject) (lispObject, error) {
	return es.stub("float-time") // Source file: timefns.c
}

func (es *emacsStubs) formatTimeString_autogen(format_string, timeval, zone lispObject) (lispObject, error) {
	return es.stub("format-time-string") // Source file: timefns.c
}

func (es *emacsStubs) setTimeZoneRule_autogen(tz lispObject) (lispObject, error) {
	return es.stub("set-time-zone-rule") // Source file: timefns.c
}

func (es *emacsStubs) timeAdd_autogen(a, b lispObject) (lispObject, error) {
	return es.stub("time-add") // Source file: timefns.c
}

func (es *emacsStubs) timeConvert_autogen(time, form lispObject) (lispObject, error) {
	return es.stub("time-convert") // Source file: timefns.c
}

func (es *emacsStubs) timeEqualP_autogen(a, b lispObject) (lispObject, error) {
	return es.stub("time-equal-p") // Source file: timefns.c
}

func (es *emacsStubs) timeLessP_autogen(a, b lispObject) (lispObject, error) {
	return es.stub("time-less-p") // Source file: timefns.c
}

func (es *emacsStubs) timeSubtract_autogen(a, b lispObject) (lispObject, error) {
	return es.stub("time-subtract") // Source file: timefns.c
}

func (es *emacsStubs) treesitAvailableP_autogen() (lispObject, error) {
	return es.stub("treesit-available-p") // Source file: treesit.c
}

func (es *emacsStubs) treesitCompiledQueryP_autogen(object lispObject) (lispObject, error) {
	return es.stub("treesit-compiled-query-p") // Source file: treesit.c
}

func (es *emacsStubs) treesitInduceSparseTree_autogen(root, predicate, process_fn, depth lispObject) (lispObject, error) {
	return es.stub("treesit-induce-sparse-tree") // Source file: treesit.c
}

func (es *emacsStubs) treesitLanguageAbiVersion_autogen(language lispObject) (lispObject, error) {
	return es.stub("treesit-language-abi-version") // Source file: treesit.c
}

func (es *emacsStubs) treesitLanguageAvailableP_autogen(language, detail lispObject) (lispObject, error) {
	return es.stub("treesit-language-available-p") // Source file: treesit.c
}

func (es *emacsStubs) treesitLibraryAbiVersion_autogen(min_compatible lispObject) (lispObject, error) {
	return es.stub("treesit-library-abi-version") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeCheck_autogen(node, property lispObject) (lispObject, error) {
	return es.stub("treesit-node-check") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeChild_autogen(node, n, named lispObject) (lispObject, error) {
	return es.stub("treesit-node-child") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeChildByFieldName_autogen(node, field_name lispObject) (lispObject, error) {
	return es.stub("treesit-node-child-by-field-name") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeChildCount_autogen(node, named lispObject) (lispObject, error) {
	return es.stub("treesit-node-child-count") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeDescendantForRange_autogen(node, beg, end, named lispObject) (lispObject, error) {
	return es.stub("treesit-node-descendant-for-range") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeEnd_autogen(node lispObject) (lispObject, error) {
	return es.stub("treesit-node-end") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeEq_autogen(node1, node2 lispObject) (lispObject, error) {
	return es.stub("treesit-node-eq") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeFieldNameForChild_autogen(node, n lispObject) (lispObject, error) {
	return es.stub("treesit-node-field-name-for-child") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeFirstChildForPos_autogen(node, pos, named lispObject) (lispObject, error) {
	return es.stub("treesit-node-first-child-for-pos") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeMatchP_autogen(node, predicate lispObject) (lispObject, error) {
	return es.stub("treesit-node-match-p") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeNextSibling_autogen(node, named lispObject) (lispObject, error) {
	return es.stub("treesit-node-next-sibling") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeP_autogen(object lispObject) (lispObject, error) {
	return es.stub("treesit-node-p") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeParent_autogen(node lispObject) (lispObject, error) {
	return es.stub("treesit-node-parent") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeParser_autogen(node lispObject) (lispObject, error) {
	return es.stub("treesit-node-parser") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodePrevSibling_autogen(node, named lispObject) (lispObject, error) {
	return es.stub("treesit-node-prev-sibling") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeStart_autogen(node lispObject) (lispObject, error) {
	return es.stub("treesit-node-start") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeString_autogen(node lispObject) (lispObject, error) {
	return es.stub("treesit-node-string") // Source file: treesit.c
}

func (es *emacsStubs) treesitNodeType_autogen(node lispObject) (lispObject, error) {
	return es.stub("treesit-node-type") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserAddNotifier_autogen(parser, function lispObject) (lispObject, error) {
	return es.stub("treesit-parser-add-notifier") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserBuffer_autogen(parser lispObject) (lispObject, error) {
	return es.stub("treesit-parser-buffer") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserCreate_autogen(language, buffer, no_reuse lispObject) (lispObject, error) {
	return es.stub("treesit-parser-create") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserDelete_autogen(parser lispObject) (lispObject, error) {
	return es.stub("treesit-parser-delete") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserIncludedRanges_autogen(parser lispObject) (lispObject, error) {
	return es.stub("treesit-parser-included-ranges") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserLanguage_autogen(parser lispObject) (lispObject, error) {
	return es.stub("treesit-parser-language") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserList_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("treesit-parser-list") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserNotifiers_autogen(parser lispObject) (lispObject, error) {
	return es.stub("treesit-parser-notifiers") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserP_autogen(object lispObject) (lispObject, error) {
	return es.stub("treesit-parser-p") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserRemoveNotifier_autogen(parser, function lispObject) (lispObject, error) {
	return es.stub("treesit-parser-remove-notifier") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserRootNode_autogen(parser lispObject) (lispObject, error) {
	return es.stub("treesit-parser-root-node") // Source file: treesit.c
}

func (es *emacsStubs) treesitParserSetIncludedRanges_autogen(parser, ranges lispObject) (lispObject, error) {
	return es.stub("treesit-parser-set-included-ranges") // Source file: treesit.c
}

func (es *emacsStubs) treesitPatternExpand_autogen(pattern lispObject) (lispObject, error) {
	return es.stub("treesit-pattern-expand") // Source file: treesit.c
}

func (es *emacsStubs) treesitQueryCapture_autogen(node, query, beg, end, node_only lispObject) (lispObject, error) {
	return es.stub("treesit-query-capture") // Source file: treesit.c
}

func (es *emacsStubs) treesitQueryCompile_autogen(language, query, eager lispObject) (lispObject, error) {
	return es.stub("treesit-query-compile") // Source file: treesit.c
}

func (es *emacsStubs) treesitQueryExpand_autogen(query lispObject) (lispObject, error) {
	return es.stub("treesit-query-expand") // Source file: treesit.c
}

func (es *emacsStubs) treesitQueryLanguage_autogen(query lispObject) (lispObject, error) {
	return es.stub("treesit-query-language") // Source file: treesit.c
}

func (es *emacsStubs) treesitQueryP_autogen(object lispObject) (lispObject, error) {
	return es.stub("treesit-query-p") // Source file: treesit.c
}

func (es *emacsStubs) treesitSearchForward_autogen(start, predicate, backward, all lispObject) (lispObject, error) {
	return es.stub("treesit-search-forward") // Source file: treesit.c
}

func (es *emacsStubs) treesitSearchSubtree_autogen(node, predicate, backward, all, depth lispObject) (lispObject, error) {
	return es.stub("treesit-search-subtree") // Source file: treesit.c
}

func (es *emacsStubs) treesitSubtreeStat_autogen(node lispObject) (lispObject, error) {
	return es.stub("treesit-subtree-stat") // Source file: treesit.c
}

func (es *emacsStubs) undoBoundary_autogen() (lispObject, error) {
	return es.stub("undo-boundary") // Source file: undo.c
}

func (es *emacsStubs) w16GetClipboardData_autogen(frame lispObject) (lispObject, error) {
	return es.stub("w16-get-clipboard-data") // Source file: w16select.c
}

func (es *emacsStubs) w16SelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return es.stub("w16-selection-exists-p") // Source file: w16select.c
}

func (es *emacsStubs) w16SetClipboardData_autogen(string, frame lispObject) (lispObject, error) {
	return es.stub("w16-set-clipboard-data") // Source file: w16select.c
}

func (es *emacsStubs) getScreenColor_autogen() (lispObject, error) {
	return es.stub("get-screen-color") // Source file: w32console.c
}

func (es *emacsStubs) setCursorSize_autogen(size lispObject) (lispObject, error) {
	return es.stub("set-cursor-size") // Source file: w32console.c
}

func (es *emacsStubs) setScreenColor_autogen(foreground, background lispObject) (lispObject, error) {
	return es.stub("set-screen-color") // Source file: w32console.c
}

func (es *emacsStubs) w32BatteryStatus_autogen() (lispObject, error) {
	return es.stub("w32-battery-status") // Source file: w32cygwinx.c
}

func (es *emacsStubs) defaultPrinterName_autogen() (lispObject, error) {
	return es.stub("default-printer-name") // Source file: w32fns.c
}

func (es *emacsStubs) setMessageBeep_autogen(sound lispObject) (lispObject, error) {
	return es.stub("set-message-beep") // Source file: w32fns.c
}

func (es *emacsStubs) systemMoveFileToTrash_autogen(filename lispObject) (lispObject, error) {
	return es.stub("system-move-file-to-trash") // Source file: w32fns.c
}

func (es *emacsStubs) w32MenuBarInUse_autogen() (lispObject, error) {
	return es.stub("w32--menu-bar-in-use") // Source file: w32fns.c
}

func (es *emacsStubs) w32DefineRgbColor_autogen(red, green, blue, name lispObject) (lispObject, error) {
	return es.stub("w32-define-rgb-color") // Source file: w32fns.c
}

func (es *emacsStubs) w32DisplayMonitorAttributesList_autogen(display lispObject) (lispObject, error) {
	return es.stub("w32-display-monitor-attributes-list") // Source file: w32fns.c
}

func (es *emacsStubs) w32FrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return es.stub("w32-frame-edges") // Source file: w32fns.c
}

func (es *emacsStubs) w32FrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return es.stub("w32-frame-geometry") // Source file: w32fns.c
}

func (es *emacsStubs) w32FrameListZOrder_autogen(display lispObject) (lispObject, error) {
	return es.stub("w32-frame-list-z-order") // Source file: w32fns.c
}

func (es *emacsStubs) w32FrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return es.stub("w32-frame-restack") // Source file: w32fns.c
}

func (es *emacsStubs) w32GetImeOpenStatus_autogen() (lispObject, error) {
	return es.stub("w32-get-ime-open-status") // Source file: w32fns.c
}

func (es *emacsStubs) w32MouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return es.stub("w32-mouse-absolute-pixel-position") // Source file: w32fns.c
}

func (es *emacsStubs) w32NotificationClose_autogen(id lispObject) (lispObject, error) {
	return es.stub("w32-notification-close") // Source file: w32fns.c
}

func (es *emacsStubs) w32NotificationNotify_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("w32-notification-notify") // Source file: w32fns.c
}

func (es *emacsStubs) w32ReadRegistry_autogen(root, key, name lispObject) (lispObject, error) {
	return es.stub("w32-read-registry") // Source file: w32fns.c
}

func (es *emacsStubs) w32ReconstructHotKey_autogen(hotkeyid lispObject) (lispObject, error) {
	return es.stub("w32-reconstruct-hot-key") // Source file: w32fns.c
}

func (es *emacsStubs) w32RegisterHotKey_autogen(key lispObject) (lispObject, error) {
	return es.stub("w32-register-hot-key") // Source file: w32fns.c
}

func (es *emacsStubs) w32RegisteredHotKeys_autogen() (lispObject, error) {
	return es.stub("w32-registered-hot-keys") // Source file: w32fns.c
}

func (es *emacsStubs) w32SendSysCommand_autogen(command, frame lispObject) (lispObject, error) {
	return es.stub("w32-send-sys-command") // Source file: w32fns.c
}

func (es *emacsStubs) w32SetImeOpenStatus_autogen(status lispObject) (lispObject, error) {
	return es.stub("w32-set-ime-open-status") // Source file: w32fns.c
}

func (es *emacsStubs) w32SetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return es.stub("w32-set-mouse-absolute-pixel-position") // Source file: w32fns.c
}

func (es *emacsStubs) w32SetWallpaper_autogen(image_file lispObject) (lispObject, error) {
	return es.stub("w32-set-wallpaper") // Source file: w32fns.c
}

func (es *emacsStubs) w32ShellExecute_autogen(operation, document, parameters, show_flag lispObject) (lispObject, error) {
	return es.stub("w32-shell-execute") // Source file: w32fns.c
}

func (es *emacsStubs) w32ToggleLockKey_autogen(key, new_state lispObject) (lispObject, error) {
	return es.stub("w32-toggle-lock-key") // Source file: w32fns.c
}

func (es *emacsStubs) w32UnregisterHotKey_autogen(key lispObject) (lispObject, error) {
	return es.stub("w32-unregister-hot-key") // Source file: w32fns.c
}

func (es *emacsStubs) w32WindowExistsP_autogen(class, name lispObject) (lispObject, error) {
	return es.stub("w32-window-exists-p") // Source file: w32fns.c
}

func (es *emacsStubs) xChangeWindowProperty_autogen(prop, value, frame, type_, format, outer_p lispObject) (lispObject, error) {
	return es.stub("x-change-window-property") // Source file: w32fns.c
}

func (es *emacsStubs) xChangeWindowProperty_1_autogen(prop, value, frame, type_, format, outer_p, window_id lispObject) (lispObject, error) {
	return es.stub("x-change-window-property") // Source file: xfns.c
}

func (es *emacsStubs) xDeleteWindowProperty_autogen(prop, frame lispObject) (lispObject, error) {
	return es.stub("x-delete-window-property") // Source file: w32fns.c
}

func (es *emacsStubs) xDeleteWindowProperty_1_autogen(prop, frame, window_id lispObject) (lispObject, error) {
	return es.stub("x-delete-window-property") // Source file: xfns.c
}

func (es *emacsStubs) xServerVendor_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-server-vendor") // Source file: w32fns.c
}

func (es *emacsStubs) xServerVendor_1_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-server-vendor") // Source file: xfns.c
}

func (es *emacsStubs) xServerVersion_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-server-version") // Source file: w32fns.c
}

func (es *emacsStubs) xServerVersion_1_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-server-version") // Source file: xfns.c
}

func (es *emacsStubs) xSynchronize_autogen(on, display lispObject) (lispObject, error) {
	return es.stub("x-synchronize") // Source file: w32fns.c
}

func (es *emacsStubs) xSynchronize_1_autogen(on, terminal lispObject) (lispObject, error) {
	return es.stub("x-synchronize") // Source file: xfns.c
}

func (es *emacsStubs) xWindowProperty_autogen(prop, frame, type_, source, delete_p, vector_ret_p lispObject) (lispObject, error) {
	return es.stub("x-window-property") // Source file: w32fns.c
}

func (es *emacsStubs) xWindowProperty_1_autogen(prop, frame, type_, window_id, delete_p, vector_ret_p lispObject) (lispObject, error) {
	return es.stub("x-window-property") // Source file: xfns.c
}

func (es *emacsStubs) w32notifyAddWatch_autogen(file, filter, callback lispObject) (lispObject, error) {
	return es.stub("w32notify-add-watch") // Source file: w32notify.c
}

func (es *emacsStubs) w32notifyRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return es.stub("w32notify-rm-watch") // Source file: w32notify.c
}

func (es *emacsStubs) w32notifyValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return es.stub("w32notify-valid-p") // Source file: w32notify.c
}

func (es *emacsStubs) w32ApplicationType_autogen(program lispObject) (lispObject, error) {
	return es.stub("w32-application-type") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetCodepageCharset_autogen(cp lispObject) (lispObject, error) {
	return es.stub("w32-get-codepage-charset") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetConsoleCodepage_autogen() (lispObject, error) {
	return es.stub("w32-get-console-codepage") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetConsoleOutputCodepage_autogen() (lispObject, error) {
	return es.stub("w32-get-console-output-codepage") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetCurrentLocaleId_autogen() (lispObject, error) {
	return es.stub("w32-get-current-locale-id") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetDefaultLocaleId_autogen(userp lispObject) (lispObject, error) {
	return es.stub("w32-get-default-locale-id") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetKeyboardLayout_autogen() (lispObject, error) {
	return es.stub("w32-get-keyboard-layout") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetLocaleInfo_autogen(lcid, longform lispObject) (lispObject, error) {
	return es.stub("w32-get-locale-info") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetValidCodepages_autogen() (lispObject, error) {
	return es.stub("w32-get-valid-codepages") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetValidKeyboardLayouts_autogen() (lispObject, error) {
	return es.stub("w32-get-valid-keyboard-layouts") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetValidLocaleIds_autogen() (lispObject, error) {
	return es.stub("w32-get-valid-locale-ids") // Source file: w32proc.c
}

func (es *emacsStubs) w32HasWinsock_autogen(load_now lispObject) (lispObject, error) {
	return es.stub("w32-has-winsock") // Source file: w32proc.c
}

func (es *emacsStubs) w32LongFileName_autogen(filename lispObject) (lispObject, error) {
	return es.stub("w32-long-file-name") // Source file: w32proc.c
}

func (es *emacsStubs) w32SetConsoleCodepage_autogen(cp lispObject) (lispObject, error) {
	return es.stub("w32-set-console-codepage") // Source file: w32proc.c
}

func (es *emacsStubs) w32SetConsoleOutputCodepage_autogen(cp lispObject) (lispObject, error) {
	return es.stub("w32-set-console-output-codepage") // Source file: w32proc.c
}

func (es *emacsStubs) w32SetCurrentLocale_autogen(lcid lispObject) (lispObject, error) {
	return es.stub("w32-set-current-locale") // Source file: w32proc.c
}

func (es *emacsStubs) w32SetKeyboardLayout_autogen(layout lispObject) (lispObject, error) {
	return es.stub("w32-set-keyboard-layout") // Source file: w32proc.c
}

func (es *emacsStubs) w32SetProcessPriority_autogen(process, priority lispObject) (lispObject, error) {
	return es.stub("w32-set-process-priority") // Source file: w32proc.c
}

func (es *emacsStubs) w32ShortFileName_autogen(filename lispObject) (lispObject, error) {
	return es.stub("w32-short-file-name") // Source file: w32proc.c
}

func (es *emacsStubs) w32UnloadWinsock_autogen() (lispObject, error) {
	return es.stub("w32-unload-winsock") // Source file: w32proc.c
}

func (es *emacsStubs) w32GetClipboardData_autogen(ignored lispObject) (lispObject, error) {
	return es.stub("w32-get-clipboard-data") // Source file: w32select.c
}

func (es *emacsStubs) w32SelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return es.stub("w32-selection-exists-p") // Source file: w32select.c
}

func (es *emacsStubs) w32SelectionTargets_autogen(selection, terminal lispObject) (lispObject, error) {
	return es.stub("w32-selection-targets") // Source file: w32select.c
}

func (es *emacsStubs) w32SetClipboardData_autogen(string, ignored lispObject) (lispObject, error) {
	return es.stub("w32-set-clipboard-data") // Source file: w32select.c
}

func (es *emacsStubs) coordinatesInWindowP_autogen(coordinates, window lispObject) (lispObject, error) {
	return es.stub("coordinates-in-window-p") // Source file: window.c
}

func (es *emacsStubs) currentWindowConfiguration_autogen(frame lispObject) (lispObject, error) {
	return es.stub("current-window-configuration") // Source file: window.c
}

func (es *emacsStubs) deleteOtherWindowsInternal_autogen(window, root lispObject) (lispObject, error) {
	return es.stub("delete-other-windows-internal") // Source file: window.c
}

func (es *emacsStubs) deleteWindowInternal_autogen(window lispObject) (lispObject, error) {
	return es.stub("delete-window-internal") // Source file: window.c
}

func (es *emacsStubs) forceWindowUpdate_autogen(object lispObject) (lispObject, error) {
	return es.stub("force-window-update") // Source file: window.c
}

func (es *emacsStubs) frameFirstWindow_autogen(frame_or_window lispObject) (lispObject, error) {
	return es.stub("frame-first-window") // Source file: window.c
}

func (es *emacsStubs) frameOldSelectedWindow_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame-old-selected-window") // Source file: window.c
}

func (es *emacsStubs) frameRootWindow_autogen(frame_or_window lispObject) (lispObject, error) {
	return es.stub("frame-root-window") // Source file: window.c
}

func (es *emacsStubs) frameSelectedWindow_autogen(frame_or_window lispObject) (lispObject, error) {
	return es.stub("frame-selected-window") // Source file: window.c
}

func (es *emacsStubs) getBufferWindow_autogen(buffer_or_name, all_frames lispObject) (lispObject, error) {
	return es.stub("get-buffer-window") // Source file: window.c
}

func (es *emacsStubs) minibufferSelectedWindow_autogen() (lispObject, error) {
	return es.stub("minibuffer-selected-window") // Source file: window.c
}

func (es *emacsStubs) minibufferWindow_autogen(frame lispObject) (lispObject, error) {
	return es.stub("minibuffer-window") // Source file: window.c
}

func (es *emacsStubs) moveToWindowLine_autogen(arg lispObject) (lispObject, error) {
	return es.stub("move-to-window-line") // Source file: window.c
}

func (es *emacsStubs) nextWindow_autogen(window, minibuf, all_frames lispObject) (lispObject, error) {
	return es.stub("next-window") // Source file: window.c
}

func (es *emacsStubs) oldSelectedWindow_autogen() (lispObject, error) {
	return es.stub("old-selected-window") // Source file: window.c
}

func (es *emacsStubs) otherWindowForScrolling_autogen() (lispObject, error) {
	return es.stub("other-window-for-scrolling") // Source file: window.c
}

func (es *emacsStubs) posVisibleInWindowP_autogen(pos, window, partially lispObject) (lispObject, error) {
	return es.stub("pos-visible-in-window-p") // Source file: window.c
}

func (es *emacsStubs) previousWindow_autogen(window, minibuf, all_frames lispObject) (lispObject, error) {
	return es.stub("previous-window") // Source file: window.c
}

func (es *emacsStubs) recenter_autogen(arg, redisplay lispObject) (lispObject, error) {
	return es.stub("recenter") // Source file: window.c
}

func (es *emacsStubs) resizeMiniWindowInternal_autogen(window lispObject) (lispObject, error) {
	return es.stub("resize-mini-window-internal") // Source file: window.c
}

func (es *emacsStubs) runWindowConfigurationChangeHook_autogen(frame lispObject) (lispObject, error) {
	return es.stub("run-window-configuration-change-hook") // Source file: window.c
}

func (es *emacsStubs) runWindowScrollFunctions_autogen(window lispObject) (lispObject, error) {
	return es.stub("run-window-scroll-functions") // Source file: window.c
}

func (es *emacsStubs) scrollDown_autogen(arg lispObject) (lispObject, error) {
	return es.stub("scroll-down") // Source file: window.c
}

func (es *emacsStubs) scrollLeft_autogen(arg, set_minimum lispObject) (lispObject, error) {
	return es.stub("scroll-left") // Source file: window.c
}

func (es *emacsStubs) scrollRight_autogen(arg, set_minimum lispObject) (lispObject, error) {
	return es.stub("scroll-right") // Source file: window.c
}

func (es *emacsStubs) scrollUp_autogen(arg lispObject) (lispObject, error) {
	return es.stub("scroll-up") // Source file: window.c
}

func (es *emacsStubs) selectWindow_autogen(window, norecord lispObject) (lispObject, error) {
	return es.stub("select-window") // Source file: window.c
}

func (es *emacsStubs) selectedWindow_autogen() (lispObject, error) {
	return es.stub("selected-window") // Source file: window.c
}

func (es *emacsStubs) setFrameSelectedWindow_autogen(frame, window, norecord lispObject) (lispObject, error) {
	return es.stub("set-frame-selected-window") // Source file: window.c
}

func (es *emacsStubs) setWindowBuffer_autogen(window, buffer_or_name, keep_margins lispObject) (lispObject, error) {
	return es.stub("set-window-buffer") // Source file: window.c
}

func (es *emacsStubs) setWindowCombinationLimit_autogen(window, limit lispObject) (lispObject, error) {
	return es.stub("set-window-combination-limit") // Source file: window.c
}

func (es *emacsStubs) setWindowConfiguration_autogen(configuration, dont_set_frame, dont_set_miniwindow lispObject) (lispObject, error) {
	return es.stub("set-window-configuration") // Source file: window.c
}

func (es *emacsStubs) setWindowDedicatedP_autogen(window, flag lispObject) (lispObject, error) {
	return es.stub("set-window-dedicated-p") // Source file: window.c
}

func (es *emacsStubs) setWindowDisplayTable_autogen(window, table lispObject) (lispObject, error) {
	return es.stub("set-window-display-table") // Source file: window.c
}

func (es *emacsStubs) setWindowFringes_autogen(window, left_width, right_width, outside_margins, persistent lispObject) (lispObject, error) {
	return es.stub("set-window-fringes") // Source file: window.c
}

func (es *emacsStubs) setWindowHscroll_autogen(window, ncol lispObject) (lispObject, error) {
	return es.stub("set-window-hscroll") // Source file: window.c
}

func (es *emacsStubs) setWindowMargins_autogen(window, left_width, right_width lispObject) (lispObject, error) {
	return es.stub("set-window-margins") // Source file: window.c
}

func (es *emacsStubs) setWindowNewNormal_autogen(window, size lispObject) (lispObject, error) {
	return es.stub("set-window-new-normal") // Source file: window.c
}

func (es *emacsStubs) setWindowNewPixel_autogen(window, size, add lispObject) (lispObject, error) {
	return es.stub("set-window-new-pixel") // Source file: window.c
}

func (es *emacsStubs) setWindowNewTotal_autogen(window, size, add lispObject) (lispObject, error) {
	return es.stub("set-window-new-total") // Source file: window.c
}

func (es *emacsStubs) setWindowNextBuffers_autogen(window, next_buffers lispObject) (lispObject, error) {
	return es.stub("set-window-next-buffers") // Source file: window.c
}

func (es *emacsStubs) setWindowParameter_autogen(window, parameter, value lispObject) (lispObject, error) {
	return es.stub("set-window-parameter") // Source file: window.c
}

func (es *emacsStubs) setWindowPoint_autogen(window, pos lispObject) (lispObject, error) {
	return es.stub("set-window-point") // Source file: window.c
}

func (es *emacsStubs) setWindowPrevBuffers_autogen(window, prev_buffers lispObject) (lispObject, error) {
	return es.stub("set-window-prev-buffers") // Source file: window.c
}

func (es *emacsStubs) setWindowScrollBars_autogen(window, width, vertical_type, height, horizontal_type, persistent lispObject) (lispObject, error) {
	return es.stub("set-window-scroll-bars") // Source file: window.c
}

func (es *emacsStubs) setWindowStart_autogen(window, pos, noforce lispObject) (lispObject, error) {
	return es.stub("set-window-start") // Source file: window.c
}

func (es *emacsStubs) setWindowVscroll_autogen(window, vscroll, pixels_p, preserve_vscroll_p lispObject) (lispObject, error) {
	return es.stub("set-window-vscroll") // Source file: window.c
}

func (es *emacsStubs) splitWindowInternal_autogen(old, pixel_size, side, normal_size lispObject) (lispObject, error) {
	return es.stub("split-window-internal") // Source file: window.c
}

func (es *emacsStubs) windowAt_autogen(x, y, frame lispObject) (lispObject, error) {
	return es.stub("window-at") // Source file: window.c
}

func (es *emacsStubs) windowBodyHeight_autogen(window, pixelwise lispObject) (lispObject, error) {
	return es.stub("window-body-height") // Source file: window.c
}

func (es *emacsStubs) windowBodyWidth_autogen(window, pixelwise lispObject) (lispObject, error) {
	return es.stub("window-body-width") // Source file: window.c
}

func (es *emacsStubs) windowBottomDividerWidth_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-bottom-divider-width") // Source file: window.c
}

func (es *emacsStubs) windowBuffer_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-buffer") // Source file: window.c
}

func (es *emacsStubs) windowBumpUseTime_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-bump-use-time") // Source file: window.c
}

func (es *emacsStubs) windowCombinationLimit_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-combination-limit") // Source file: window.c
}

func (es *emacsStubs) windowConfigurationEqualP_autogen(x, y lispObject) (lispObject, error) {
	return es.stub("window-configuration-equal-p") // Source file: window.c
}

func (es *emacsStubs) windowConfigurationFrame_autogen(config lispObject) (lispObject, error) {
	return es.stub("window-configuration-frame") // Source file: window.c
}

func (es *emacsStubs) windowConfigurationP_autogen(object lispObject) (lispObject, error) {
	return es.stub("window-configuration-p") // Source file: window.c
}

func (es *emacsStubs) windowDedicatedP_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-dedicated-p") // Source file: window.c
}

func (es *emacsStubs) windowDisplayTable_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-display-table") // Source file: window.c
}

func (es *emacsStubs) windowEnd_autogen(window, update lispObject) (lispObject, error) {
	return es.stub("window-end") // Source file: window.c
}

func (es *emacsStubs) windowFrame_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-frame") // Source file: window.c
}

func (es *emacsStubs) windowFringes_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-fringes") // Source file: window.c
}

func (es *emacsStubs) windowHeaderLineHeight_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-header-line-height") // Source file: window.c
}

func (es *emacsStubs) windowHscroll_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-hscroll") // Source file: window.c
}

func (es *emacsStubs) windowLeftChild_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-left-child") // Source file: window.c
}

func (es *emacsStubs) windowLeftColumn_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-left-column") // Source file: window.c
}

func (es *emacsStubs) windowLineHeight_autogen(line, window lispObject) (lispObject, error) {
	return es.stub("window-line-height") // Source file: window.c
}

func (es *emacsStubs) windowLinesPixelDimensions_autogen(window, first, last, body, inverse, left lispObject) (lispObject, error) {
	return es.stub("window-lines-pixel-dimensions") // Source file: window.c
}

func (es *emacsStubs) windowList_autogen(frame, minibuf, window lispObject) (lispObject, error) {
	return es.stub("window-list") // Source file: window.c
}

func (es *emacsStubs) windowList1_autogen(window, minibuf, all_frames lispObject) (lispObject, error) {
	return es.stub("window-list-1") // Source file: window.c
}

func (es *emacsStubs) windowLiveP_autogen(object lispObject) (lispObject, error) {
	return es.stub("window-live-p") // Source file: window.c
}

func (es *emacsStubs) windowMargins_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-margins") // Source file: window.c
}

func (es *emacsStubs) windowMinibufferP_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-minibuffer-p") // Source file: window.c
}

func (es *emacsStubs) windowModeLineHeight_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-mode-line-height") // Source file: window.c
}

func (es *emacsStubs) windowNewNormal_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-new-normal") // Source file: window.c
}

func (es *emacsStubs) windowNewPixel_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-new-pixel") // Source file: window.c
}

func (es *emacsStubs) windowNewTotal_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-new-total") // Source file: window.c
}

func (es *emacsStubs) windowNextBuffers_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-next-buffers") // Source file: window.c
}

func (es *emacsStubs) windowNextSibling_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-next-sibling") // Source file: window.c
}

func (es *emacsStubs) windowNormalSize_autogen(window, horizontal lispObject) (lispObject, error) {
	return es.stub("window-normal-size") // Source file: window.c
}

func (es *emacsStubs) windowOldBodyPixelHeight_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-old-body-pixel-height") // Source file: window.c
}

func (es *emacsStubs) windowOldBodyPixelWidth_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-old-body-pixel-width") // Source file: window.c
}

func (es *emacsStubs) windowOldBuffer_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-old-buffer") // Source file: window.c
}

func (es *emacsStubs) windowOldPixelHeight_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-old-pixel-height") // Source file: window.c
}

func (es *emacsStubs) windowOldPixelWidth_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-old-pixel-width") // Source file: window.c
}

func (es *emacsStubs) windowOldPoint_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-old-point") // Source file: window.c
}

func (es *emacsStubs) windowParameter_autogen(window, parameter lispObject) (lispObject, error) {
	return es.stub("window-parameter") // Source file: window.c
}

func (es *emacsStubs) windowParameters_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-parameters") // Source file: window.c
}

func (es *emacsStubs) windowParent_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-parent") // Source file: window.c
}

func (es *emacsStubs) windowPixelHeight_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-pixel-height") // Source file: window.c
}

func (es *emacsStubs) windowPixelLeft_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-pixel-left") // Source file: window.c
}

func (es *emacsStubs) windowPixelTop_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-pixel-top") // Source file: window.c
}

func (es *emacsStubs) windowPixelWidth_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-pixel-width") // Source file: window.c
}

func (es *emacsStubs) windowPoint_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-point") // Source file: window.c
}

func (es *emacsStubs) windowPrevBuffers_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-prev-buffers") // Source file: window.c
}

func (es *emacsStubs) windowPrevSibling_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-prev-sibling") // Source file: window.c
}

func (es *emacsStubs) windowResizeApply_autogen(frame, horizontal lispObject) (lispObject, error) {
	return es.stub("window-resize-apply") // Source file: window.c
}

func (es *emacsStubs) windowResizeApplyTotal_autogen(frame, horizontal lispObject) (lispObject, error) {
	return es.stub("window-resize-apply-total") // Source file: window.c
}

func (es *emacsStubs) windowRightDividerWidth_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-right-divider-width") // Source file: window.c
}

func (es *emacsStubs) windowScrollBarHeight_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-scroll-bar-height") // Source file: window.c
}

func (es *emacsStubs) windowScrollBarWidth_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-scroll-bar-width") // Source file: window.c
}

func (es *emacsStubs) windowScrollBars_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-scroll-bars") // Source file: window.c
}

func (es *emacsStubs) windowStart_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-start") // Source file: window.c
}

func (es *emacsStubs) windowTabLineHeight_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-tab-line-height") // Source file: window.c
}

func (es *emacsStubs) windowTextHeight_autogen(window, pixelwise lispObject) (lispObject, error) {
	return es.stub("window-text-height") // Source file: window.c
}

func (es *emacsStubs) windowTextWidth_autogen(window, pixelwise lispObject) (lispObject, error) {
	return es.stub("window-text-width") // Source file: window.c
}

func (es *emacsStubs) windowTopChild_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-top-child") // Source file: window.c
}

func (es *emacsStubs) windowTopLine_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-top-line") // Source file: window.c
}

func (es *emacsStubs) windowTotalHeight_autogen(window, round lispObject) (lispObject, error) {
	return es.stub("window-total-height") // Source file: window.c
}

func (es *emacsStubs) windowTotalWidth_autogen(window, round lispObject) (lispObject, error) {
	return es.stub("window-total-width") // Source file: window.c
}

func (es *emacsStubs) windowUseTime_autogen(window lispObject) (lispObject, error) {
	return es.stub("window-use-time") // Source file: window.c
}

func (es *emacsStubs) windowValidP_autogen(object lispObject) (lispObject, error) {
	return es.stub("window-valid-p") // Source file: window.c
}

func (es *emacsStubs) windowVscroll_autogen(window, pixels_p lispObject) (lispObject, error) {
	return es.stub("window-vscroll") // Source file: window.c
}

func (es *emacsStubs) windowp_autogen(object lispObject) (lispObject, error) {
	return es.stub("windowp") // Source file: window.c
}

func (es *emacsStubs) bidiFindOverriddenDirectionality_autogen(from, to, object, base_dir lispObject) (lispObject, error) {
	return es.stub("bidi-find-overridden-directionality") // Source file: xdisp.c
}

func (es *emacsStubs) bidiResolvedLevels_autogen(vpos lispObject) (lispObject, error) {
	return es.stub("bidi-resolved-levels") // Source file: xdisp.c
}

func (es *emacsStubs) bufferTextPixelSize_autogen(buffer_or_name, window, x_limit, y_limit lispObject) (lispObject, error) {
	return es.stub("buffer-text-pixel-size") // Source file: xdisp.c
}

func (es *emacsStubs) currentBidiParagraphDirection_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("current-bidi-paragraph-direction") // Source file: xdisp.c
}

func (es *emacsStubs) displayLineIsContinuedP_autogen() (lispObject, error) {
	return es.stub("display--line-is-continued-p") // Source file: xdisp.c
}

func (es *emacsStubs) dumpFrameGlyphMatrix_autogen() (lispObject, error) {
	return es.stub("dump-frame-glyph-matrix") // Source file: xdisp.c
}

func (es *emacsStubs) dumpGlyphMatrix_autogen(glyphs lispObject) (lispObject, error) {
	return es.stub("dump-glyph-matrix") // Source file: xdisp.c
}

func (es *emacsStubs) dumpGlyphRow_autogen(row, glyphs lispObject) (lispObject, error) {
	return es.stub("dump-glyph-row") // Source file: xdisp.c
}

func (es *emacsStubs) dumpTabBarRow_autogen(row, glyphs lispObject) (lispObject, error) {
	return es.stub("dump-tab-bar-row") // Source file: xdisp.c
}

func (es *emacsStubs) dumpToolBarRow_autogen(row, glyphs lispObject) (lispObject, error) {
	return es.stub("dump-tool-bar-row") // Source file: xdisp.c
}

func (es *emacsStubs) formatModeLine_autogen(format, face, window, buffer lispObject) (lispObject, error) {
	return es.stub("format-mode-line") // Source file: xdisp.c
}

func (es *emacsStubs) getDisplayProperty_autogen(position, prop, object, properties lispObject) (lispObject, error) {
	return es.stub("get-display-property") // Source file: xdisp.c
}

func (es *emacsStubs) invisibleP_autogen(pos lispObject) (lispObject, error) {
	return es.stub("invisible-p") // Source file: xdisp.c
}

func (es *emacsStubs) linePixelHeight_autogen() (lispObject, error) {
	return es.stub("line-pixel-height") // Source file: xdisp.c
}

func (es *emacsStubs) longLineOptimizationsP_autogen() (lispObject, error) {
	return es.stub("long-line-optimizations-p") // Source file: xdisp.c
}

func (es *emacsStubs) lookupImageMap_autogen(map_, x, y lispObject) (lispObject, error) {
	return es.stub("lookup-image-map") // Source file: xdisp.c
}

func (es *emacsStubs) movePointVisually_autogen(direction lispObject) (lispObject, error) {
	return es.stub("move-point-visually") // Source file: xdisp.c
}

func (es *emacsStubs) setBufferRedisplay_autogen(symbol, newval, op, where lispObject) (lispObject, error) {
	return es.stub("set-buffer-redisplay") // Source file: xdisp.c
}

func (es *emacsStubs) tabBarHeight_autogen(frame, pixelwise lispObject) (lispObject, error) {
	return es.stub("tab-bar-height") // Source file: xdisp.c
}

func (es *emacsStubs) toolBarHeight_autogen(frame, pixelwise lispObject) (lispObject, error) {
	return es.stub("tool-bar-height") // Source file: xdisp.c
}

func (es *emacsStubs) traceRedisplay_autogen(arg lispObject) (lispObject, error) {
	return es.stub("trace-redisplay") // Source file: xdisp.c
}

func (es *emacsStubs) traceToStderr_autogen(args ...lispObject) (lispObject, error) {
	return es.stub("trace-to-stderr") // Source file: xdisp.c
}

func (es *emacsStubs) windowTextPixelSize_autogen(window, from, to, x_limit, y_limit, mode_lines, ignore_line_at_end lispObject) (lispObject, error) {
	return es.stub("window-text-pixel-size") // Source file: xdisp.c
}

func (es *emacsStubs) bitmapSpecP_autogen(object lispObject) (lispObject, error) {
	return es.stub("bitmap-spec-p") // Source file: xfaces.c
}

func (es *emacsStubs) clearFaceCache_autogen(thoroughly lispObject) (lispObject, error) {
	return es.stub("clear-face-cache") // Source file: xfaces.c
}

func (es *emacsStubs) colorDistance_autogen(color1, color2, frame, metric lispObject) (lispObject, error) {
	return es.stub("color-distance") // Source file: xfaces.c
}

func (es *emacsStubs) colorGrayP_autogen(color, frame lispObject) (lispObject, error) {
	return es.stub("color-gray-p") // Source file: xfaces.c
}

func (es *emacsStubs) colorSupportedP_autogen(color, frame, background_p lispObject) (lispObject, error) {
	return es.stub("color-supported-p") // Source file: xfaces.c
}

func (es *emacsStubs) colorValuesFromColorSpec_autogen(spec lispObject) (lispObject, error) {
	return es.stub("color-values-from-color-spec") // Source file: xfaces.c
}

func (es *emacsStubs) displaySupportsFaceAttributesP_autogen(attributes, display lispObject) (lispObject, error) {
	return es.stub("display-supports-face-attributes-p") // Source file: xfaces.c
}

func (es *emacsStubs) dumpColors_autogen() (lispObject, error) {
	return es.stub("dump-colors") // Source file: xfaces.c
}

func (es *emacsStubs) dumpFace_autogen(n lispObject) (lispObject, error) {
	return es.stub("dump-face") // Source file: xfaces.c
}

func (es *emacsStubs) faceAttributeRelativeP_autogen(attribute, value lispObject) (lispObject, error) {
	return es.stub("face-attribute-relative-p") // Source file: xfaces.c, attributes: const
}

func (es *emacsStubs) faceAttributesAsVector_autogen(plist lispObject) (lispObject, error) {
	return es.stub("face-attributes-as-vector") // Source file: xfaces.c
}

func (es *emacsStubs) faceFont_autogen(face, frame, character lispObject) (lispObject, error) {
	return es.stub("face-font") // Source file: xfaces.c
}

func (es *emacsStubs) frameFaceHashTable_autogen(frame lispObject) (lispObject, error) {
	return es.stub("frame--face-hash-table") // Source file: xfaces.c
}

func (es *emacsStubs) internalCopyLispFace_autogen(from, to, frame, new_frame lispObject) (lispObject, error) {
	return es.stub("internal-copy-lisp-face") // Source file: xfaces.c
}

func (es *emacsStubs) internalFaceXGetResource_autogen(resource, class, frame lispObject) (lispObject, error) {
	return es.stub("internal-face-x-get-resource") // Source file: xfaces.c
}

func (es *emacsStubs) internalGetLispFaceAttribute_autogen(symbol, keyword, frame lispObject) (lispObject, error) {
	return es.stub("internal-get-lisp-face-attribute") // Source file: xfaces.c
}

func (es *emacsStubs) internalLispFaceAttributeValues_autogen(attr lispObject) (lispObject, error) {
	return es.stub("internal-lisp-face-attribute-values") // Source file: xfaces.c
}

func (es *emacsStubs) internalLispFaceEmptyP_autogen(face, frame lispObject) (lispObject, error) {
	return es.stub("internal-lisp-face-empty-p") // Source file: xfaces.c
}

func (es *emacsStubs) internalLispFaceEqualP_autogen(face1, face2, frame lispObject) (lispObject, error) {
	return es.stub("internal-lisp-face-equal-p") // Source file: xfaces.c
}

func (es *emacsStubs) internalLispFaceP_autogen(face, frame lispObject) (lispObject, error) {
	return es.stub("internal-lisp-face-p") // Source file: xfaces.c
}

func (es *emacsStubs) internalMakeLispFace_autogen(face, frame lispObject) (lispObject, error) {
	return es.stub("internal-make-lisp-face") // Source file: xfaces.c
}

func (es *emacsStubs) internalMergeInGlobalFace_autogen(face, frame lispObject) (lispObject, error) {
	return es.stub("internal-merge-in-global-face") // Source file: xfaces.c
}

func (es *emacsStubs) internalSetAlternativeFontFamilyAlist_autogen(alist lispObject) (lispObject, error) {
	return es.stub("internal-set-alternative-font-family-alist") // Source file: xfaces.c
}

func (es *emacsStubs) internalSetAlternativeFontRegistryAlist_autogen(alist lispObject) (lispObject, error) {
	return es.stub("internal-set-alternative-font-registry-alist") // Source file: xfaces.c
}

func (es *emacsStubs) internalSetFontSelectionOrder_autogen(order lispObject) (lispObject, error) {
	return es.stub("internal-set-font-selection-order") // Source file: xfaces.c
}

func (es *emacsStubs) internalSetLispFaceAttribute_autogen(face, attr, value, frame lispObject) (lispObject, error) {
	return es.stub("internal-set-lisp-face-attribute") // Source file: xfaces.c
}

func (es *emacsStubs) internalSetLispFaceAttributeFromResource_autogen(face, attr, value, frame lispObject) (lispObject, error) {
	return es.stub("internal-set-lisp-face-attribute-from-resource") // Source file: xfaces.c
}

func (es *emacsStubs) mergeFaceAttribute_autogen(attribute, value1, value2 lispObject) (lispObject, error) {
	return es.stub("merge-face-attribute") // Source file: xfaces.c
}

func (es *emacsStubs) showFaceResources_autogen() (lispObject, error) {
	return es.stub("show-face-resources") // Source file: xfaces.c
}

func (es *emacsStubs) ttySuppressBoldInverseDefaultColors_autogen(suppress lispObject) (lispObject, error) {
	return es.stub("tty-suppress-bold-inverse-default-colors") // Source file: xfaces.c
}

func (es *emacsStubs) xFamilyFonts_autogen(family, frame lispObject) (lispObject, error) {
	return es.stub("x-family-fonts") // Source file: xfaces.c
}

func (es *emacsStubs) xListFonts_autogen(pattern, face, frame, maximum, width lispObject) (lispObject, error) {
	return es.stub("x-list-fonts") // Source file: xfaces.c
}

func (es *emacsStubs) xLoadColorFile_autogen(filename lispObject) (lispObject, error) {
	return es.stub("x-load-color-file") // Source file: xfaces.c
}

func (es *emacsStubs) xBackspaceDeleteKeysP_autogen(frame lispObject) (lispObject, error) {
	return es.stub("x-backspace-delete-keys-p") // Source file: xfns.c
}

func (es *emacsStubs) xBeginDrag_autogen(targets, action, frame, return_frame, allow_current_frame, follow_tooltip lispObject) (lispObject, error) {
	return es.stub("x-begin-drag") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayMonitorAttributesList_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-display-monitor-attributes-list") // Source file: xfns.c
}

func (es *emacsStubs) xDisplayLastUserTime_autogen(time_object, terminal lispObject) (lispObject, error) {
	return es.stub("x-display-set-last-user-time") // Source file: xfns.c
}

func (es *emacsStubs) xDoubleBufferedP_autogen(frame lispObject) (lispObject, error) {
	return es.stub("x-double-buffered-p") // Source file: xfns.c
}

func (es *emacsStubs) xFrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return es.stub("x-frame-edges") // Source file: xfns.c
}

func (es *emacsStubs) xFrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return es.stub("x-frame-geometry") // Source file: xfns.c
}

func (es *emacsStubs) xFrameListZOrder_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-frame-list-z-order") // Source file: xfns.c
}

func (es *emacsStubs) xFrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return es.stub("x-frame-restack") // Source file: xfns.c
}

func (es *emacsStubs) xGetModifierMasks_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-get-modifier-masks") // Source file: xfns.c
}

func (es *emacsStubs) xGetPageSetup_autogen() (lispObject, error) {
	return es.stub("x-get-page-setup") // Source file: xfns.c
}

func (es *emacsStubs) xInternalFocusInputContext_autogen(focus lispObject) (lispObject, error) {
	return es.stub("x-internal-focus-input-context") // Source file: xfns.c
}

func (es *emacsStubs) xMouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return es.stub("x-mouse-absolute-pixel-position") // Source file: xfns.c
}

func (es *emacsStubs) xPageSetupDialog_autogen() (lispObject, error) {
	return es.stub("x-page-setup-dialog") // Source file: xfns.c
}

func (es *emacsStubs) xPrintFramesDialog_autogen(frames lispObject) (lispObject, error) {
	return es.stub("x-print-frames-dialog") // Source file: xfns.c
}

func (es *emacsStubs) xServerInputExtensionVersion_autogen(terminal lispObject) (lispObject, error) {
	return es.stub("x-server-input-extension-version") // Source file: xfns.c
}

func (es *emacsStubs) xSetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return es.stub("x-set-mouse-absolute-pixel-position") // Source file: xfns.c
}

func (es *emacsStubs) xTestStringConversion_autogen(frame, position, direction, operation, factor lispObject) (lispObject, error) {
	return es.stub("x-test-string-conversion") // Source file: xfns.c
}

func (es *emacsStubs) xTranslateCoordinates_autogen(frame, source_window, dest_window, source_x, source_y, require_child lispObject) (lispObject, error) {
	return es.stub("x-translate-coordinates") // Source file: xfns.c
}

func (es *emacsStubs) xUsesOldGtkDialog_autogen() (lispObject, error) {
	return es.stub("x-uses-old-gtk-dialog") // Source file: xfns.c
}

func (es *emacsStubs) xWindowPropertyAttributes_autogen(prop, frame, window_id lispObject) (lispObject, error) {
	return es.stub("x-window-property-attributes") // Source file: xfns.c
}

func (es *emacsStubs) xWmSetSizeHint_autogen(frame lispObject) (lispObject, error) {
	return es.stub("x-wm-set-size-hint") // Source file: xfns.c
}

func (es *emacsStubs) libxmlAvailableP_autogen() (lispObject, error) {
	return es.stub("libxml-available-p") // Source file: xml.c
}

func (es *emacsStubs) libxmlParseHtmlRegion_autogen(start, end, base_url, discard_comments lispObject) (lispObject, error) {
	return es.stub("libxml-parse-html-region") // Source file: xml.c
}

func (es *emacsStubs) libxmlParseXmlRegion_autogen(start, end, base_url, discard_comments lispObject) (lispObject, error) {
	return es.stub("libxml-parse-xml-region") // Source file: xml.c
}

func (es *emacsStubs) xDisownSelectionInternal_autogen(selection, time_object, terminal lispObject) (lispObject, error) {
	return es.stub("x-disown-selection-internal") // Source file: xselect.c
}

func (es *emacsStubs) xGetAtomName_autogen(value, frame lispObject) (lispObject, error) {
	return es.stub("x-get-atom-name") // Source file: xselect.c
}

func (es *emacsStubs) xGetLocalSelection_autogen(value, target lispObject) (lispObject, error) {
	return es.stub("x-get-local-selection") // Source file: xselect.c
}

func (es *emacsStubs) xGetSelectionInternal_autogen(selection_symbol, target_type, time_stamp, terminal lispObject) (lispObject, error) {
	return es.stub("x-get-selection-internal") // Source file: xselect.c
}

func (es *emacsStubs) xOwnSelectionInternal_autogen(selection, value, frame lispObject) (lispObject, error) {
	return es.stub("x-own-selection-internal") // Source file: xselect.c
}

func (es *emacsStubs) xRegisterDndAtom_autogen(atom, frame lispObject) (lispObject, error) {
	return es.stub("x-register-dnd-atom") // Source file: xselect.c
}

func (es *emacsStubs) xSelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return es.stub("x-selection-exists-p") // Source file: xselect.c
}

func (es *emacsStubs) xSelectionOwnerP_autogen(selection, terminal lispObject) (lispObject, error) {
	return es.stub("x-selection-owner-p") // Source file: xselect.c
}

func (es *emacsStubs) xSendClientMessage_autogen(display, dest, from, message_type, format, values lispObject) (lispObject, error) {
	return es.stub("x-send-client-message") // Source file: xselect.c
}

func (es *emacsStubs) fontGetSystemFont_autogen() (lispObject, error) {
	return es.stub("font-get-system-font") // Source file: xsettings.c
}

func (es *emacsStubs) fontGetSystemNormalFont_autogen() (lispObject, error) {
	return es.stub("font-get-system-normal-font") // Source file: xsettings.c
}

func (es *emacsStubs) toolBarGetSystemStyle_autogen() (lispObject, error) {
	return es.stub("tool-bar-get-system-style") // Source file: xsettings.c
}

func (es *emacsStubs) handleSaveSession_autogen(event lispObject) (lispObject, error) {
	return es.stub("handle-save-session") // Source file: xsmfns.c
}

func (es *emacsStubs) deleteXwidgetView_autogen(xwidget_view lispObject) (lispObject, error) {
	return es.stub("delete-xwidget-view") // Source file: xwidget.c
}

func (es *emacsStubs) getBufferXwidgets_autogen(buffer lispObject) (lispObject, error) {
	return es.stub("get-buffer-xwidgets") // Source file: xwidget.c
}

func (es *emacsStubs) killXwidget_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("kill-xwidget") // Source file: xwidget.c
}

func (es *emacsStubs) makeXwidget_autogen(type_, title, width, height, arguments, buffer, related lispObject) (lispObject, error) {
	return es.stub("make-xwidget") // Source file: xwidget.c
}

func (es *emacsStubs) setXwidgetBuffer_autogen(xwidget, buffer lispObject) (lispObject, error) {
	return es.stub("set-xwidget-buffer") // Source file: xwidget.c
}

func (es *emacsStubs) setXwidgetPlist_autogen(xwidget, plist lispObject) (lispObject, error) {
	return es.stub("set-xwidget-plist") // Source file: xwidget.c
}

func (es *emacsStubs) setXwidgetQueryOnExitFlag_autogen(xwidget, flag lispObject) (lispObject, error) {
	return es.stub("set-xwidget-query-on-exit-flag") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetBuffer_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-buffer") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetInfo_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-info") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetLiveP_autogen(object lispObject) (lispObject, error) {
	return es.stub("xwidget-live-p") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetPerformLispyEvent_autogen(xwidget, event, frame lispObject) (lispObject, error) {
	return es.stub("xwidget-perform-lispy-event") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetPlist_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-plist") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetQueryOnExitFlag_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-query-on-exit-flag") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetResize_autogen(xwidget, new_width, new_height lispObject) (lispObject, error) {
	return es.stub("xwidget-resize") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetSizeRequest_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-size-request") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetViewInfo_autogen(xwidget_view lispObject) (lispObject, error) {
	return es.stub("xwidget-view-info") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetViewLookup_autogen(xwidget, window lispObject) (lispObject, error) {
	return es.stub("xwidget-view-lookup") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetViewModel_autogen(xwidget_view lispObject) (lispObject, error) {
	return es.stub("xwidget-view-model") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetViewP_autogen(object lispObject) (lispObject, error) {
	return es.stub("xwidget-view-p") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetViewWindow_autogen(xwidget_view lispObject) (lispObject, error) {
	return es.stub("xwidget-view-window") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitBackForwardList_autogen(xwidget, limit lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-back-forward-list") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitEstimatedLoadProgress_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-estimated-load-progress") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitExecuteScript_autogen(xwidget, script, fun lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-execute-script") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitFinishSearch_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-finish-search") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitGotoHistory_autogen(xwidget, rel_pos lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-goto-history") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitGotoUri_autogen(xwidget, uri lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-goto-uri") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitLoadHtml_autogen(xwidget, text, base_uri lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-load-html") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitNextResult_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-next-result") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitPreviousResult_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-previous-result") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitSearch_autogen(query, xwidget, case_insensitive, backwards, wrap_around lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-search") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitSetCookieStorageFile_autogen(xwidget, file lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-set-cookie-storage-file") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitStopLoading_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-stop-loading") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitTitle_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-title") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitUri_autogen(xwidget lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-uri") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetWebkitZoom_autogen(xwidget, factor lispObject) (lispObject, error) {
	return es.stub("xwidget-webkit-zoom") // Source file: xwidget.c
}

func (es *emacsStubs) xwidgetp_autogen(object lispObject) (lispObject, error) {
	return es.stub("xwidgetp") // Source file: xwidget.c
}

func (ec *execContext) symbolsOfEmacs_autogen() {
	es := &emacsStubs{ec: ec}
	ec.stubs = es
	ec.defSubrM(nil, "bool-vector", es.boolVector_autogen, 0)
	ec.defSubr2(nil, "cons", es.cons_autogen, 2)
	ec.defSubr0(nil, "garbage-collect", es.garbageCollect_autogen)
	ec.defSubr1(nil, "garbage-collect-maybe", es.garbageCollectMaybe_autogen, 1)
	ec.defSubrM(nil, "list", es.list_autogen, 0)
	ec.defSubr2(nil, "make-bool-vector", es.makeBoolVector_autogen, 2)
	ec.defSubrM(nil, "make-byte-code", es.makeByteCode_autogen, 4)
	ec.defSubrM(nil, "make-closure", es.makeClosure_autogen, 1)
	ec.defSubr1(nil, "make-finalizer", es.makeFinalizer_autogen, 1)
	ec.defSubr2(nil, "make-list", es.makeList_autogen, 2)
	ec.defSubr0(nil, "make-marker", es.makeMarker_autogen)
	ec.defSubr3(nil, "make-record", es.makeRecord_autogen, 3)
	ec.defSubr3(nil, "make-string", es.makeString_autogen, 2)
	ec.defSubr1(nil, "make-symbol", es.makeSymbol_autogen, 1)
	ec.defSubr2(nil, "make-vector", es.makeVector_autogen, 2)
	ec.defSubr0(nil, "malloc-info", es.mallocInfo_autogen)
	ec.defSubr1(nil, "malloc-trim", es.mallocTrim_autogen, 0)
	ec.defSubr0(nil, "memory-info", es.memoryInfo_autogen)
	ec.defSubr0(nil, "memory-use-counts", es.memoryUseCounts_autogen)
	ec.defSubr1(nil, "purecopy", es.purecopy_autogen, 1)
	ec.defSubrM(nil, "record", es.record_autogen, 1)
	ec.defSubr1(nil, "suspicious-object", es.suspiciousObject_autogen, 1)
	ec.defSubrM(nil, "vector", es.vector_autogen, 0)
	ec.defSubr0(nil, "debug-timer-check", es.debugTimerCheck_autogen)
	ec.defSubr1(nil, "barf-if-buffer-read-only", es.barfIfBufferReadOnly_autogen, 0)
	ec.defSubr1(nil, "buffer-base-buffer", es.bufferBaseBuffer_autogen, 0)
	ec.defSubr1(nil, "buffer-chars-modified-tick", es.bufferCharsModifiedTick_autogen, 0)
	ec.defSubr1(nil, "buffer-enable-undo", es.bufferEnableUndo_autogen, 0)
	ec.defSubr1(nil, "buffer-file-name", es.bufferFileName_autogen, 0)
	ec.defSubr1(nil, "buffer-list", es.bufferList_autogen, 0)
	ec.defSubr1(nil, "buffer-live-p", es.bufferLiveP_autogen, 1)
	ec.defSubr2(nil, "buffer-local-value", es.bufferLocalValue_autogen, 2)
	ec.defSubr1(nil, "buffer-local-variables", es.bufferLocalVariables_autogen, 0)
	ec.defSubr1(nil, "buffer-modified-p", es.bufferModifiedP_autogen, 0)
	ec.defSubr1(nil, "buffer-modified-tick", es.bufferModifiedTick_autogen, 0)
	ec.defSubr1(nil, "buffer-name", es.bufferName_autogen, 0)
	ec.defSubr1(nil, "buffer-swap-text", es.bufferSwapText_autogen, 1)
	ec.defSubr1(nil, "bury-buffer-internal", es.buryBufferInternal_autogen, 1)
	ec.defSubr0(nil, "current-buffer", es.currentBuffer_autogen)
	ec.defSubr1(nil, "delete-all-overlays", es.deleteAllOverlays_autogen, 0)
	ec.defSubr1(nil, "delete-overlay", es.deleteOverlay_autogen, 1)
	ec.defSubr0(nil, "erase-buffer", es.eraseBuffer_autogen)
	ec.defSubr1(nil, "force-mode-line-update", es.forceModeLineUpdate_autogen, 0)
	ec.defSubr2(nil, "generate-new-buffer-name", es.generateNewBufferName_autogen, 1)
	ec.defSubr1(nil, "get-buffer", es.getBuffer_autogen, 1)
	ec.defSubr2(nil, "get-buffer-create", es.getBufferCreate_autogen, 1)
	ec.defSubr1(nil, "get-file-buffer", es.getFileBuffer_autogen, 1)
	ec.defSubr2(nil, "internal--set-buffer-modified-tick", es.internalSetBufferModifiedTick_autogen, 1)
	ec.defSubr1(nil, "kill-all-local-variables", es.killAllLocalVariables_autogen, 0)
	ec.defSubr1(nil, "kill-buffer", es.killBuffer_autogen, 0)
	ec.defSubr4(nil, "make-indirect-buffer", es.makeIndirectBuffer_autogen, 2)
	ec.defSubr5(nil, "make-overlay", es.makeOverlay_autogen, 2)
	ec.defSubr4(nil, "move-overlay", es.moveOverlay_autogen, 3)
	ec.defSubr1(nil, "next-overlay-change", es.nextOverlayChange_autogen, 1)
	ec.defSubr3(nil, "other-buffer", es.otherBuffer_autogen, 0)
	ec.defSubr1(nil, "overlay-buffer", es.overlayBuffer_autogen, 1)
	ec.defSubr1(nil, "overlay-end", es.overlayEnd_autogen, 1)
	ec.defSubr2(nil, "overlay-get", es.overlayGet_autogen, 2)
	ec.defSubr0(nil, "overlay-lists", es.overlayLists_autogen)
	ec.defSubr1(nil, "overlay-properties", es.overlayProperties_autogen, 1)
	ec.defSubr3(nil, "overlay-put", es.overlayPut_autogen, 3)
	ec.defSubr1(nil, "overlay-recenter", es.overlayRecenter_autogen, 1)
	ec.defSubr1(nil, "overlay-start", es.overlayStart_autogen, 1)
	ec.defSubr1(nil, "overlay-tree", es.overlayTree_autogen, 0)
	ec.defSubr1(nil, "overlayp", es.overlayp_autogen, 1)
	ec.defSubr2(nil, "overlays-at", es.overlaysAt_autogen, 1)
	ec.defSubr2(nil, "overlays-in", es.overlaysIn_autogen, 2)
	ec.defSubr1(nil, "previous-overlay-change", es.previousOverlayChange_autogen, 1)
	ec.defSubr2(nil, "rename-buffer", es.renameBuffer_autogen, 1)
	ec.defSubr1(nil, "restore-buffer-modified-p", es.restoreBufferModifiedP_autogen, 1)
	ec.defSubr1(nil, "set-buffer", es.setBuffer_autogen, 1)
	ec.defSubr1(nil, "set-buffer-major-mode", es.setBufferMajorMode_autogen, 1)
	ec.defSubr1(nil, "set-buffer-modified-p", es.setBufferModifiedP_autogen, 1)
	ec.defSubr1(nil, "set-buffer-multibyte", es.setBufferMultibyte_autogen, 1)
	ec.defSubr3(nil, "byte-code", es.byteCode_autogen, 3)
	ec.defSubr0(nil, "internal-stack-stats", es.internalStackStats_autogen)
	ec.defSubr3(nil, "call-interactively", es.callInteractively_autogen, 1)
	ec.defSubrM(nil, "funcall-interactively", es.funcallInteractively_autogen, 1)
	ec.defSubrU(nil, "interactive", es.interactive_autogen, 0)
	ec.defSubr1(nil, "prefix-numeric-value", es.prefixNumericValue_autogen, 1)
	ec.defSubrM(nil, "call-process", es.callProcess_autogen, 1)
	ec.defSubrM(nil, "call-process-region", es.callProcessRegion_autogen, 3)
	ec.defSubr2(nil, "getenv-internal", es.getenvInternal_autogen, 1)
	ec.defSubr1(nil, "capitalize", es.capitalize_autogen, 1)
	ec.defSubr3(nil, "capitalize-region", es.capitalizeRegion_autogen, 2)
	ec.defSubr1(nil, "capitalize-word", es.capitalizeWord_autogen, 1)
	ec.defSubr1(nil, "downcase", es.downcase_autogen, 1)
	ec.defSubr3(nil, "downcase-region", es.downcaseRegion_autogen, 2)
	ec.defSubr1(nil, "downcase-word", es.downcaseWord_autogen, 1)
	ec.defSubr1(nil, "upcase", es.upcase_autogen, 1)
	ec.defSubr1(nil, "upcase-initials", es.upcaseInitials_autogen, 1)
	ec.defSubr3(nil, "upcase-initials-region", es.upcaseInitialsRegion_autogen, 2)
	ec.defSubr3(nil, "upcase-region", es.upcaseRegion_autogen, 2)
	ec.defSubr1(nil, "upcase-word", es.upcaseWord_autogen, 1)
	ec.defSubr1(nil, "case-table-p", es.caseTableP_autogen, 1)
	ec.defSubr0(nil, "current-case-table", es.currentCaseTable_autogen)
	ec.defSubr1(nil, "set-case-table", es.setCaseTable_autogen, 1)
	ec.defSubr1(nil, "set-standard-case-table", es.setStandardCaseTable_autogen, 1)
	ec.defSubr0(nil, "standard-case-table", es.standardCaseTable_autogen)
	ec.defSubr2(nil, "category-docstring", es.categoryDocstring_autogen, 1)
	ec.defSubr1(nil, "category-set-mnemonics", es.categorySetMnemonics_autogen, 1)
	ec.defSubr0(nil, "category-table", es.categoryTable_autogen)
	ec.defSubr1(nil, "category-table-p", es.categoryTableP_autogen, 1)
	ec.defSubr1(nil, "char-category-set", es.charCategorySet_autogen, 1)
	ec.defSubr1(nil, "copy-category-table", es.copyCategoryTable_autogen, 0)
	ec.defSubr3(nil, "define-category", es.defineCategory_autogen, 2)
	ec.defSubr1(nil, "get-unused-category", es.getUnusedCategory_autogen, 0)
	ec.defSubr1(nil, "make-category-set", es.makeCategorySet_autogen, 1)
	ec.defSubr0(nil, "make-category-table", es.makeCategoryTable_autogen)
	ec.defSubr4(nil, "modify-category-entry", es.modifyCategoryEntry_autogen, 2)
	ec.defSubr1(nil, "set-category-table", es.setCategoryTable_autogen, 1)
	ec.defSubr0(nil, "standard-category-table", es.standardCategoryTable_autogen)
	ec.defSubr2(nil, "ccl-execute", es.cclExecute_autogen, 2)
	ec.defSubr5(nil, "ccl-execute-on-string", es.cclExecuteOnString_autogen, 3)
	ec.defSubr1(nil, "ccl-program-p", es.cclProgramP_autogen, 1)
	ec.defSubr2(nil, "register-ccl-program", es.registerCclProgram_autogen, 2)
	ec.defSubr2(nil, "register-code-conversion-map", es.registerCodeConversionMap_autogen, 2)
	ec.defSubr1(nil, "char-resolve-modifiers", es.charResolveModifiers_autogen, 1)
	ec.defSubr1(nil, "char-width", es.charWidth_autogen, 1)
	ec.defSubr2(nil, "characterp", es.characterp_autogen, 1)
	ec.defSubr2(nil, "get-byte", es.getByte_autogen, 0)
	ec.defSubr1(nil, "max-char", es.maxChar_autogen, 0)
	ec.defSubr1(nil, "multibyte-char-to-unibyte", es.multibyteCharToUnibyte_autogen, 1)
	ec.defSubrM(nil, "string", es.string_autogen, 0)
	ec.defSubr3(nil, "string-width", es.stringWidth_autogen, 1)
	ec.defSubr1(nil, "unibyte-char-to-multibyte", es.unibyteCharToMultibyte_autogen, 1)
	ec.defSubrM(nil, "unibyte-string", es.unibyteString_autogen, 0)
	ec.defSubr2(nil, "char-charset", es.charCharset_autogen, 1)
	ec.defSubr1(nil, "charset-after", es.charsetAfter_autogen, 0)
	ec.defSubr1(nil, "charset-id-internal", es.charsetIdInternal_autogen, 0)
	ec.defSubr1(nil, "charset-plist", es.charsetPlist_autogen, 1)
	ec.defSubr1(nil, "charset-priority-list", es.charsetPriorityList_autogen, 0)
	ec.defSubr1(nil, "charsetp", es.charsetp_autogen, 1)
	ec.defSubr0(nil, "clear-charset-maps", es.clearCharsetMaps_autogen)
	ec.defSubr4(nil, "declare-equiv-charset", es.declareEquivCharset_autogen, 4)
	ec.defSubr2(nil, "decode-char", es.decodeChar_autogen, 2)
	ec.defSubr2(nil, "define-charset-alias", es.defineCharsetAlias_autogen, 2)
	ec.defSubrM(nil, "define-charset-internal", es.defineCharsetInternal_autogen, emacsConstant_charset_arg_max_autogen)
	ec.defSubr2(nil, "encode-char", es.encodeChar_autogen, 2)
	ec.defSubr3(nil, "find-charset-region", es.findCharsetRegion_autogen, 2)
	ec.defSubr2(nil, "find-charset-string", es.findCharsetString_autogen, 1)
	ec.defSubr2(nil, "get-unused-iso-final-char", es.getUnusedIsoFinalChar_autogen, 2)
	ec.defSubr3(nil, "iso-charset", es.isoCharset_autogen, 3)
	ec.defSubr5(nil, "make-char", es.makeChar_autogen, 1)
	ec.defSubr5(nil, "map-charset-chars", es.mapCharsetChars_autogen, 2)
	ec.defSubr2(nil, "set-charset-plist", es.setCharsetPlist_autogen, 2)
	ec.defSubrM(nil, "set-charset-priority", es.setCharsetPriority_autogen, 1)
	ec.defSubr1(nil, "sort-charsets", es.sortCharsets_autogen, 1)
	ec.defSubr1(nil, "split-char", es.splitChar_autogen, 1)
	ec.defSubr3(nil, "unify-charset", es.unifyCharset_autogen, 1)
	ec.defSubr2(nil, "char-table-extra-slot", es.charTableExtraSlot_autogen, 2)
	ec.defSubr1(nil, "char-table-parent", es.charTableParent_autogen, 1)
	ec.defSubr2(nil, "char-table-range", es.charTableRange_autogen, 2)
	ec.defSubr1(nil, "char-table-subtype", es.charTableSubtype_autogen, 1)
	ec.defSubr2(nil, "get-unicode-property-internal", es.getUnicodePropertyInternal_autogen, 2)
	ec.defSubr2(nil, "make-char-table", es.makeCharTable_autogen, 1)
	ec.defSubr2(nil, "map-char-table", es.mapCharTable_autogen, 2)
	ec.defSubr2(nil, "optimize-char-table", es.optimizeCharTable_autogen, 1)
	ec.defSubr3(nil, "put-unicode-property-internal", es.putUnicodePropertyInternal_autogen, 3)
	ec.defSubr3(nil, "set-char-table-extra-slot", es.setCharTableExtraSlot_autogen, 3)
	ec.defSubr2(nil, "set-char-table-parent", es.setCharTableParent_autogen, 2)
	ec.defSubr3(nil, "set-char-table-range", es.setCharTableRange_autogen, 3)
	ec.defSubr1(nil, "unicode-property-table-internal", es.unicodePropertyTableInternal_autogen, 1)
	ec.defSubr1(nil, "backward-char", es.backwardChar_autogen, 0)
	ec.defSubr1(nil, "beginning-of-line", es.beginningOfLine_autogen, 0)
	ec.defSubr2(nil, "delete-char", es.deleteChar_autogen, 1)
	ec.defSubr1(nil, "end-of-line", es.endOfLine_autogen, 0)
	ec.defSubr1(nil, "forward-char", es.forwardChar_autogen, 0)
	ec.defSubr1(nil, "forward-line", es.forwardLine_autogen, 0)
	ec.defSubr2(nil, "self-insert-command", es.selfInsertCommand_autogen, 1)
	ec.defSubr1(nil, "check-coding-system", es.checkCodingSystem_autogen, 1)
	ec.defSubr3(nil, "check-coding-systems-region", es.checkCodingSystemsRegion_autogen, 3)
	ec.defSubr1(nil, "coding-system-aliases", es.codingSystemAliases_autogen, 1)
	ec.defSubr1(nil, "coding-system-base", es.codingSystemBase_autogen, 1)
	ec.defSubr1(nil, "coding-system-eol-type", es.codingSystemEolType_autogen, 1)
	ec.defSubr1(nil, "coding-system-p", es.codingSystemP_autogen, 1)
	ec.defSubr1(nil, "coding-system-plist", es.codingSystemPlist_autogen, 1)
	ec.defSubr1(nil, "coding-system-priority-list", es.codingSystemPriorityList_autogen, 0)
	ec.defSubr3(nil, "coding-system-put", es.codingSystemPut_autogen, 3)
	ec.defSubr1(nil, "decode-big5-char", es.decodeBig5Char_autogen, 1)
	ec.defSubr4(nil, "decode-coding-region", es.decodeCodingRegion_autogen, 3)
	ec.defSubr4(nil, "decode-coding-string", es.decodeCodingString_autogen, 2)
	ec.defSubr1(nil, "decode-sjis-char", es.decodeSjisChar_autogen, 1)
	ec.defSubr2(nil, "define-coding-system-alias", es.defineCodingSystemAlias_autogen, 2)
	ec.defSubrM(nil, "define-coding-system-internal", es.defineCodingSystemInternal_autogen, emacsConstant_coding_arg_max_autogen)
	ec.defSubr3(nil, "detect-coding-region", es.detectCodingRegion_autogen, 2)
	ec.defSubr2(nil, "detect-coding-string", es.detectCodingString_autogen, 1)
	ec.defSubr1(nil, "encode-big5-char", es.encodeBig5Char_autogen, 1)
	ec.defSubr4(nil, "encode-coding-region", es.encodeCodingRegion_autogen, 3)
	ec.defSubr4(nil, "encode-coding-string", es.encodeCodingString_autogen, 2)
	ec.defSubr1(nil, "encode-sjis-char", es.encodeSjisChar_autogen, 1)
	ec.defSubr3(nil, "find-coding-systems-region-internal", es.findCodingSystemsRegionInternal_autogen, 2)
	ec.defSubrM(nil, "find-operation-coding-system", es.findOperationCodingSystem_autogen, 1)
	ec.defSubr7(nil, "internal-decode-string-utf-8", es.internalDecodeStringUtf8_autogen, 7)
	ec.defSubr7(nil, "internal-encode-string-utf-8", es.internalEncodeStringUtf8_autogen, 7)
	ec.defSubr1(nil, "keyboard-coding-system", es.keyboardCodingSystem_autogen, 0)
	ec.defSubr2(nil, "read-coding-system", es.readCodingSystem_autogen, 1)
	ec.defSubr1(nil, "read-non-nil-coding-system", es.readNonNilCodingSystem_autogen, 1)
	ec.defSubrM(nil, "set-coding-system-priority", es.setCodingSystemPriority_autogen, 0)
	ec.defSubr2(nil, "set-keyboard-coding-system-internal", es.setKeyboardCodingSystemInternal_autogen, 1)
	ec.defSubr1(nil, "set-safe-terminal-coding-system-internal", es.setSafeTerminalCodingSystemInternal_autogen, 1)
	ec.defSubr2(nil, "set-terminal-coding-system-internal", es.setTerminalCodingSystemInternal_autogen, 1)
	ec.defSubr1(nil, "terminal-coding-system", es.terminalCodingSystem_autogen, 0)
	ec.defSubr5(nil, "unencodable-char-position", es.unencodableCharPosition_autogen, 3)
	ec.defSubr1(nil, "comp--compile-ctxt-to-file", es.compCompileCtxtToFile_autogen, 1)
	ec.defSubr0(nil, "comp--init-ctxt", es.compInitCtxt_autogen)
	ec.defSubr2(nil, "comp--install-trampoline", es.compInstallTrampoline_autogen, 2)
	ec.defSubr7(nil, "comp--late-register-subr", es.compLateRegisterSubr_autogen, 7)
	ec.defSubr7(nil, "comp--register-lambda", es.compRegisterLambda_autogen, 7)
	ec.defSubr7(nil, "comp--register-subr", es.compRegisterSubr_autogen, 7)
	ec.defSubr0(nil, "comp--release-ctxt", es.compReleaseCtxt_autogen)
	ec.defSubr1(nil, "comp--subr-signature", es.compSubrSignature_autogen, 1)
	ec.defSubr2(nil, "comp-el-to-eln-filename", es.compElToElnFilename_autogen, 1)
	ec.defSubr1(nil, "comp-el-to-eln-rel-filename", es.compElToElnRelFilename_autogen, 1)
	ec.defSubr0(nil, "comp-libgccjit-version", es.compLibgccjitVersion_autogen)
	ec.defSubr0(nil, "comp-native-compiler-options-effective-p", es.compNativeCompilerOptionsEffectiveP_autogen)
	ec.defSubr0(nil, "comp-native-driver-options-effective-p", es.compNativeDriverOptionsEffectiveP_autogen)
	ec.defSubr0(nil, "native-comp-available-p", es.nativeCompAvailableP_autogen)
	ec.defSubr2(nil, "native-elisp-load", es.nativeElispLoad_autogen, 1)
	ec.defSubr0(nil, "clear-composition-cache", es.clearCompositionCache_autogen)
	ec.defSubr4(nil, "compose-region-internal", es.composeRegionInternal_autogen, 2)
	ec.defSubr5(nil, "compose-string-internal", es.composeStringInternal_autogen, 3)
	ec.defSubr4(nil, "composition-get-gstring", es.compositionGetGstring_autogen, 4)
	ec.defSubr1(nil, "composition-sort-rules", es.compositionSortRules_autogen, 1)
	ec.defSubr4(nil, "find-composition-internal", es.findCompositionInternal_autogen, 4)
	ec.defSubr2(nil, "cygwin-convert-file-name-from-windows", es.cygwinConvertFileNameFromWindows_autogen, 1)
	ec.defSubr2(nil, "cygwin-convert-file-name-to-windows", es.cygwinConvertFileNameToWindows_autogen, 1)
	ec.defSubr2(nil, "%", es.rem_autogen, 2)
	ec.defSubrM(nil, "*", es.times_autogen, 0)
	ec.defSubrM(nil, "+", es.plus_autogen, 0)
	ec.defSubrM(nil, "-", es.minus_autogen, 0)
	ec.defSubrM(nil, "/", es.quo_autogen, 1)
	ec.defSubr2(nil, "/=", es.neq_autogen, 2)
	ec.defSubr1(nil, "1+", es.add1_autogen, 1)
	ec.defSubr1(nil, "1-", es.sub1_autogen, 1)
	ec.defSubrM(nil, "<", es.lss_autogen, 1)
	ec.defSubrM(nil, "<=", es.leq_autogen, 1)
	ec.defSubrM(nil, "=", es.eqlsign_autogen, 1)
	ec.defSubrM(nil, ">", es.gtr_autogen, 1)
	ec.defSubrM(nil, ">=", es.geq_autogen, 1)
	ec.defSubr2(nil, "add-variable-watcher", es.addVariableWatcher_autogen, 2)
	ec.defSubr2(nil, "aref", es.aref_autogen, 2)
	ec.defSubr1(nil, "arrayp", es.arrayp_autogen, 1)
	ec.defSubr3(nil, "aset", es.aset_autogen, 3)
	ec.defSubr2(nil, "ash", es.ash_autogen, 2)
	ec.defSubr1(nil, "atom", es.atom_autogen, 1)
	ec.defSubr1(nil, "bare-symbol", es.bareSymbol_autogen, 1)
	ec.defSubr1(nil, "bare-symbol-p", es.bareSymbolP_autogen, 1)
	ec.defSubr3(nil, "bool-vector-count-consecutive", es.boolVectorCountConsecutive_autogen, 3)
	ec.defSubr1(nil, "bool-vector-count-population", es.boolVectorCountPopulation_autogen, 1)
	ec.defSubr3(nil, "bool-vector-exclusive-or", es.boolVectorExclusiveOr_autogen, 2)
	ec.defSubr3(nil, "bool-vector-intersection", es.boolVectorIntersection_autogen, 2)
	ec.defSubr2(nil, "bool-vector-not", es.boolVectorNot_autogen, 1)
	ec.defSubr1(nil, "bool-vector-p", es.boolVectorP_autogen, 1)
	ec.defSubr3(nil, "bool-vector-set-difference", es.boolVectorSetDifference_autogen, 2)
	ec.defSubr2(nil, "bool-vector-subsetp", es.boolVectorSubsetp_autogen, 2)
	ec.defSubr3(nil, "bool-vector-union", es.boolVectorUnion_autogen, 2)
	ec.defSubr1(nil, "boundp", es.boundp_autogen, 1)
	ec.defSubr1(nil, "bufferp", es.bufferp_autogen, 1)
	ec.defSubr1(nil, "byte-code-function-p", es.byteCodeFunctionP_autogen, 1)
	ec.defSubr0(nil, "byteorder", es.byteorder_autogen)
	ec.defSubr1(nil, "car", es.car_autogen, 1)
	ec.defSubr1(nil, "car-safe", es.carSafe_autogen, 1)
	ec.defSubr1(nil, "cdr", es.cdr_autogen, 1)
	ec.defSubr1(nil, "cdr-safe", es.cdrSafe_autogen, 1)
	ec.defSubr1(nil, "char-or-string-p", es.charOrStringP_autogen, 1)
	ec.defSubr1(nil, "char-table-p", es.charTableP_autogen, 1)
	ec.defSubr1(nil, "command-modes", es.commandModes_autogen, 1)
	ec.defSubr1(nil, "condition-variable-p", es.conditionVariableP_autogen, 1)
	ec.defSubr1(nil, "consp", es.consp_autogen, 1)
	ec.defSubr3(nil, "defalias", es.defalias_autogen, 2)
	ec.defSubr1(nil, "default-boundp", es.defaultBoundp_autogen, 1)
	ec.defSubr1(nil, "default-value", es.defaultValue_autogen, 1)
	ec.defSubr2(nil, "eq", es.eq_autogen, 2)
	ec.defSubr1(nil, "fboundp", es.fboundp_autogen, 1)
	ec.defSubr1(nil, "floatp", es.floatp_autogen, 1)
	ec.defSubr1(nil, "fmakunbound", es.fmakunbound_autogen, 1)
	ec.defSubr2(nil, "fset", es.fset_autogen, 2)
	ec.defSubr1(nil, "get-variable-watchers", es.getVariableWatchers_autogen, 1)
	ec.defSubr2(nil, "indirect-function", es.indirectFunction_autogen, 1)
	ec.defSubr1(nil, "indirect-variable", es.indirectVariable_autogen, 1)
	ec.defSubr1(nil, "integer-or-marker-p", es.integerOrMarkerP_autogen, 1)
	ec.defSubr1(nil, "integerp", es.integerp_autogen, 1)
	ec.defSubr1(nil, "interactive-form", es.interactiveForm_autogen, 1)
	ec.defSubr1(nil, "keywordp", es.keywordp_autogen, 1)
	ec.defSubr1(nil, "kill-local-variable", es.killLocalVariable_autogen, 1)
	ec.defSubr1(nil, "listp", es.listp_autogen, 1)
	ec.defSubr2(nil, "local-variable-if-set-p", es.localVariableIfSetP_autogen, 1)
	ec.defSubr2(nil, "local-variable-p", es.localVariableP_autogen, 1)
	ec.defSubrM(nil, "logand", es.logand_autogen, 0)
	ec.defSubr1(nil, "logcount", es.logcount_autogen, 1)
	ec.defSubrM(nil, "logior", es.logior_autogen, 0)
	ec.defSubr1(nil, "lognot", es.lognot_autogen, 1)
	ec.defSubrM(nil, "logxor", es.logxor_autogen, 0)
	ec.defSubr1(nil, "make-local-variable", es.makeLocalVariable_autogen, 1)
	ec.defSubr1(nil, "make-variable-buffer-local", es.makeVariableBufferLocal_autogen, 1)
	ec.defSubr1(nil, "makunbound", es.makunbound_autogen, 1)
	ec.defSubr1(nil, "markerp", es.markerp_autogen, 1)
	ec.defSubrM(nil, "max", es.max_autogen, 1)
	ec.defSubrM(nil, "min", es.min_autogen, 1)
	ec.defSubr2(nil, "mod", es.mod_autogen, 2)
	ec.defSubr1(nil, "module-function-p", es.moduleFunctionP_autogen, 1)
	ec.defSubr1(nil, "multibyte-string-p", es.multibyteStringP_autogen, 1)
	ec.defSubr1(nil, "mutexp", es.mutexp_autogen, 1)
	ec.defSubr1(nil, "native-comp-unit-file", es.nativeCompUnitFile_autogen, 1)
	ec.defSubr2(nil, "native-comp-unit-set-file", es.nativeCompUnitSetFile_autogen, 2)
	ec.defSubr1(nil, "natnump", es.natnump_autogen, 1)
	ec.defSubr1(nil, "nlistp", es.nlistp_autogen, 1)
	ec.defSubr1(nil, "null", es.null_autogen, 1)
	ec.defSubr1(nil, "number-or-marker-p", es.numberOrMarkerP_autogen, 1)
	ec.defSubr1(nil, "number-to-string", es.numberToString_autogen, 1)
	ec.defSubr1(nil, "numberp", es.numberp_autogen, 1)
	ec.defSubr2(nil, "position-symbol", es.positionSymbol_autogen, 2)
	ec.defSubr1(nil, "recordp", es.recordp_autogen, 1)
	ec.defSubr1(nil, "remove-pos-from-symbol", es.removePosFromSymbol_autogen, 1)
	ec.defSubr2(nil, "remove-variable-watcher", es.removeVariableWatcher_autogen, 2)
	ec.defSubr1(nil, "sequencep", es.sequencep_autogen, 1)
	ec.defSubr2(nil, "set", es.set_autogen, 2)
	ec.defSubr2(nil, "set-default", es.setDefault_autogen, 2)
	ec.defSubr2(nil, "setcar", es.setcar_autogen, 2)
	ec.defSubr2(nil, "setcdr", es.setcdr_autogen, 2)
	ec.defSubr2(nil, "setplist", es.setplist_autogen, 2)
	ec.defSubr2(nil, "string-to-number", es.stringToNumber_autogen, 1)
	ec.defSubr1(nil, "stringp", es.stringp_autogen, 1)
	ec.defSubr1(nil, "subr-arity", es.subrArity_autogen, 1)
	ec.defSubr1(nil, "subr-name", es.subrName_autogen, 1)
	ec.defSubr1(nil, "subr-native-comp-unit", es.subrNativeCompUnit_autogen, 1)
	ec.defSubr1(nil, "subr-native-elisp-p", es.subrNativeElispP_autogen, 1)
	ec.defSubr1(nil, "subr-native-lambda-list", es.subrNativeLambdaList_autogen, 1)
	ec.defSubr1(nil, "subr-type", es.subrType_autogen, 1)
	ec.defSubr1(nil, "subrp", es.subrp_autogen, 1)
	ec.defSubr1(nil, "symbol-function", es.symbolFunction_autogen, 1)
	ec.defSubr1(nil, "symbol-name", es.symbolName_autogen, 1)
	ec.defSubr1(nil, "symbol-plist", es.symbolPlist_autogen, 1)
	ec.defSubr1(nil, "symbol-value", es.symbolValue_autogen, 1)
	ec.defSubr1(nil, "symbol-with-pos-p", es.symbolWithPosP_autogen, 1)
	ec.defSubr1(nil, "symbol-with-pos-pos", es.symbolWithPosPos_autogen, 1)
	ec.defSubr1(nil, "symbolp", es.symbolp_autogen, 1)
	ec.defSubr1(nil, "threadp", es.threadp_autogen, 1)
	ec.defSubr1(nil, "type-of", es.typeOf_autogen, 1)
	ec.defSubr1(nil, "user-ptrp", es.userPtrp_autogen, 1)
	ec.defSubr1(nil, "variable-binding-locus", es.variableBindingLocus_autogen, 1)
	ec.defSubr1(nil, "vector-or-char-table-p", es.vectorOrCharTableP_autogen, 1)
	ec.defSubr1(nil, "vectorp", es.vectorp_autogen, 1)
	ec.defSubr2(nil, "dbus--init-bus", es.dbusInitBus_autogen, 1)
	ec.defSubr1(nil, "dbus-get-unique-name", es.dbusGetUniqueName_autogen, 1)
	ec.defSubrM(nil, "dbus-message-internal", es.dbusMessageInternal_autogen, 4)
	ec.defSubr0(nil, "zlib-available-p", es.zlibAvailableP_autogen)
	ec.defSubr3(nil, "zlib-decompress-region", es.zlibDecompressRegion_autogen, 2)
	ec.defSubr5(nil, "directory-files", es.directoryFiles_autogen, 1)
	ec.defSubr6(nil, "directory-files-and-attributes", es.directoryFilesAndAttributes_autogen, 1)
	ec.defSubr2(nil, "file-attributes", es.fileAttributes_autogen, 1)
	ec.defSubr2(nil, "file-attributes-lessp", es.fileAttributesLessp_autogen, 2)
	ec.defSubr2(nil, "file-name-all-completions", es.fileNameAllCompletions_autogen, 2)
	ec.defSubr3(nil, "file-name-completion", es.fileNameCompletion_autogen, 2)
	ec.defSubr0(nil, "system-groups", es.systemGroups_autogen)
	ec.defSubr0(nil, "system-users", es.systemUsers_autogen)
	ec.defSubr1(nil, "ding", es.ding_autogen, 0)
	ec.defSubr2(nil, "display--update-for-mouse-movement", es.displayUpdateForMouseMovement_autogen, 2)
	ec.defSubr0(nil, "dump-redisplay-history", es.dumpRedisplayHistory_autogen)
	ec.defSubr1(nil, "frame-or-buffer-changed-p", es.frameOrBufferChangedP_autogen, 0)
	ec.defSubr2(nil, "internal-show-cursor", es.internalShowCursor_autogen, 2)
	ec.defSubr1(nil, "internal-show-cursor-p", es.internalShowCursorP_autogen, 0)
	ec.defSubr1(nil, "open-termscript", es.openTermscript_autogen, 1)
	ec.defSubr1(nil, "redisplay", es.redisplay_autogen, 0)
	ec.defSubr0(nil, "redraw-display", es.redrawDisplay_autogen)
	ec.defSubr1(nil, "redraw-frame", es.redrawFrame_autogen, 0)
	ec.defSubr2(nil, "send-string-to-terminal", es.sendStringToTerminal_autogen, 1)
	ec.defSubr2(nil, "sleep-for", es.sleepFor_autogen, 1)
	ec.defSubr1(nil, "Snarf-documentation", es.snarfDocumentation_autogen, 1)
	ec.defSubr2(nil, "documentation", es.documentation_autogen, 1)
	ec.defSubr3(nil, "documentation-property", es.documentationProperty_autogen, 2)
	ec.defSubr1(nil, "internal-subr-documentation", es.subrDocumentation_autogen, 1)
	ec.defSubr0(nil, "text-quoting-style", es.textQuotingStyle_autogen)
	ec.defSubr0(nil, "bobp", es.bobp_autogen)
	ec.defSubr0(nil, "bolp", es.bolp_autogen)
	ec.defSubr1(nil, "buffer-size", es.bufferSize_autogen, 0)
	ec.defSubr0(nil, "buffer-string", es.bufferString_autogen)
	ec.defSubr2(nil, "buffer-substring", es.bufferSubstring_autogen, 2)
	ec.defSubr2(nil, "buffer-substring-no-properties", es.bufferSubstringNoProperties_autogen, 2)
	ec.defSubr1(nil, "byte-to-position", es.byteToPosition_autogen, 1)
	ec.defSubr1(nil, "byte-to-string", es.byteToString_autogen, 1)
	ec.defSubr1(nil, "char-after", es.charAfter_autogen, 0)
	ec.defSubr1(nil, "char-before", es.charBefore_autogen, 0)
	ec.defSubr2(nil, "char-equal", es.charEqual_autogen, 2)
	ec.defSubr1(nil, "char-to-string", es.charToString_autogen, 1)
	ec.defSubr6(nil, "compare-buffer-substrings", es.compareBufferSubstrings_autogen, 6)
	ec.defSubr5(nil, "constrain-to-field", es.constrainToField_autogen, 2)
	ec.defSubr0(nil, "current-message", es.currentMessage_autogen)
	ec.defSubr2(nil, "delete-and-extract-region", es.deleteAndExtractRegion_autogen, 2)
	ec.defSubr1(nil, "delete-field", es.deleteField_autogen, 0)
	ec.defSubr2(nil, "delete-region", es.deleteRegion_autogen, 2)
	ec.defSubr0(nil, "emacs-pid", es.emacsPid_autogen)
	ec.defSubr0(nil, "eobp", es.eobp_autogen)
	ec.defSubr0(nil, "eolp", es.eolp_autogen)
	ec.defSubr3(nil, "field-beginning", es.fieldBeginning_autogen, 0)
	ec.defSubr3(nil, "field-end", es.fieldEnd_autogen, 0)
	ec.defSubr1(nil, "field-string", es.fieldString_autogen, 0)
	ec.defSubr1(nil, "field-string-no-properties", es.fieldStringNoProperties_autogen, 0)
	ec.defSubr0(nil, "following-char", es.followingChar_autogen)
	ec.defSubrM(nil, "format", es.format_autogen, 1)
	ec.defSubrM(nil, "format-message", es.formatMessage_autogen, 1)
	ec.defSubr0(nil, "gap-position", es.gapPosition_autogen)
	ec.defSubr0(nil, "gap-size", es.gapSize_autogen)
	ec.defSubr3(nil, "get-pos-property", es.getPosProperty_autogen, 2)
	ec.defSubr1(nil, "goto-char", es.gotoChar_autogen, 1)
	ec.defSubr0(nil, "group-gid", es.groupGid_autogen)
	ec.defSubr1(nil, "group-name", es.groupName_autogen, 1)
	ec.defSubr0(nil, "group-real-gid", es.groupRealGid_autogen)
	ec.defSubrM(nil, "insert", es.insert_autogen, 0)
	ec.defSubrM(nil, "insert-and-inherit", es.insertAndInherit_autogen, 0)
	ec.defSubrM(nil, "insert-before-markers", es.insertBeforeMarkers_autogen, 0)
	ec.defSubrM(nil, "insert-before-markers-and-inherit", es.insertAndInheritBeforeMarkers_autogen, 0)
	ec.defSubr3(nil, "insert-buffer-substring", es.insertBufferSubstring_autogen, 1)
	ec.defSubr3(nil, "insert-byte", es.insertByte_autogen, 2)
	ec.defSubr3(nil, "insert-char", es.insertChar_autogen, 1)
	ec.defSubr3(nil, "internal--labeled-narrow-to-region", es.internalLabeledNarrowToRegion_autogen, 3)
	ec.defSubr1(nil, "internal--labeled-widen", es.internalLabeledWiden_autogen, 1)
	ec.defSubr1(nil, "line-beginning-position", es.lineBeginningPosition_autogen, 0)
	ec.defSubr1(nil, "line-end-position", es.lineEndPosition_autogen, 0)
	ec.defSubr0(nil, "mark-marker", es.markMarker_autogen)
	ec.defSubrM(nil, "message", es.message_autogen, 1)
	ec.defSubrM(nil, "message-box", es.messageBox_autogen, 1)
	ec.defSubrM(nil, "message-or-box", es.messageOrBox_autogen, 1)
	ec.defSubr2(nil, "narrow-to-region", es.narrowToRegion_autogen, 2)
	ec.defSubr3(nil, "ngettext", es.ngettext_autogen, 3)
	ec.defSubr0(nil, "point", es.point_autogen)
	ec.defSubr0(nil, "point-marker", es.pointMarker_autogen)
	ec.defSubr0(nil, "point-max", es.pointMax_autogen)
	ec.defSubr0(nil, "point-max-marker", es.pointMaxMarker_autogen)
	ec.defSubr0(nil, "point-min", es.pointMin_autogen)
	ec.defSubr0(nil, "point-min-marker", es.pointMinMarker_autogen)
	ec.defSubr1(nil, "pos-bol", es.posBol_autogen, 0)
	ec.defSubr1(nil, "pos-eol", es.posEol_autogen, 0)
	ec.defSubr1(nil, "position-bytes", es.positionBytes_autogen, 1)
	ec.defSubr0(nil, "preceding-char", es.previousChar_autogen)
	ec.defSubrM(nil, "propertize", es.propertize_autogen, 1)
	ec.defSubr0(nil, "region-beginning", es.regionBeginning_autogen)
	ec.defSubr0(nil, "region-end", es.regionEnd_autogen)
	ec.defSubr3(nil, "replace-buffer-contents", es.replaceBufferContents_autogen, 1)
	ec.defSubrU(nil, "save-current-buffer", es.saveCurrentBuffer_autogen, 0)
	ec.defSubrU(nil, "save-excursion", es.saveExcursion_autogen, 0)
	ec.defSubrU(nil, "save-restriction", es.saveRestriction_autogen, 0)
	ec.defSubr1(nil, "string-to-char", es.stringToChar_autogen, 1)
	ec.defSubr5(nil, "subst-char-in-region", es.substCharInRegion_autogen, 4)
	ec.defSubr0(nil, "system-name", es.systemName_autogen)
	ec.defSubr3(nil, "translate-region-internal", es.translateRegionInternal_autogen, 3)
	ec.defSubr5(nil, "transpose-regions", es.transposeRegions_autogen, 4)
	ec.defSubr1(nil, "user-full-name", es.userFullName_autogen, 0)
	ec.defSubr1(nil, "user-login-name", es.userLoginName_autogen, 0)
	ec.defSubr0(nil, "user-real-login-name", es.userRealLoginName_autogen)
	ec.defSubr0(nil, "user-real-uid", es.userRealUid_autogen)
	ec.defSubr0(nil, "user-uid", es.userUid_autogen)
	ec.defSubr0(nil, "widen", es.widen_autogen)
	ec.defSubr1(nil, "module-load", es.moduleLoad_autogen, 1)
	ec.defSubr0(nil, "daemon-initialized", es.daemonInitialized_autogen)
	ec.defSubr0(nil, "daemonp", es.daemonp_autogen)
	ec.defSubr2(nil, "dump-emacs", es.dumpEmacs_autogen, 2)
	ec.defSubr0(nil, "invocation-directory", es.invocationDirectory_autogen)
	ec.defSubr0(nil, "invocation-name", es.invocationName_autogen)
	ec.defSubr2(nil, "kill-emacs", es.killEmacs_autogen, 0)
	ec.defSubrU(nil, "and", es.and_autogen, 0)
	ec.defSubrM(nil, "apply", es.apply_autogen, 1)
	ec.defSubr5(nil, "autoload", es.autoload_autogen, 2)
	ec.defSubr3(nil, "autoload-do-load", es.autoloadDoLoad_autogen, 1)
	ec.defSubr1(nil, "backtrace--frames-from-thread", es.backtraceFramesFromThread_autogen, 1)
	ec.defSubr2(nil, "backtrace--locals", es.backtraceLocals_autogen, 1)
	ec.defSubr2(nil, "backtrace-debug", es.backtraceDebug_autogen, 2)
	ec.defSubr3(nil, "backtrace-eval", es.backtraceEval_autogen, 2)
	ec.defSubr3(nil, "backtrace-frame--internal", es.backtraceFrameInternal_autogen, 3)
	ec.defSubrU(nil, "catch", es.catch_autogen, 1)
	ec.defSubr2(nil, "commandp", es.commandp_autogen, 1)
	ec.defSubrU(nil, "cond", es.cond_autogen, 0)
	ec.defSubrU(nil, "condition-case", es.conditionCase_autogen, 2)
	ec.defSubr1(nil, "default-toplevel-value", es.defaultToplevelValue_autogen, 1)
	ec.defSubrU(nil, "defconst", es.defconst_autogen, 2)
	ec.defSubr3(nil, "defconst-1", es.defconst1_autogen, 2)
	ec.defSubrU(nil, "defvar", es.defvar_autogen, 1)
	ec.defSubr3(nil, "defvar-1", es.defvar1_autogen, 2)
	ec.defSubr3(nil, "defvaralias", es.defvaralias_autogen, 2)
	ec.defSubr2(nil, "eval", es.eval_autogen, 1)
	ec.defSubr1(nil, "fetch-bytecode", es.fetchBytecode_autogen, 1)
	ec.defSubr1(nil, "func-arity", es.funcArity_autogen, 1)
	ec.defSubrM(nil, "funcall", es.funcall_autogen, 1)
	ec.defSubr3(nil, "funcall-with-delayed-message", es.funcallWithDelayedMessage_autogen, 3)
	ec.defSubrU(nil, "function", es.function_autogen, 1)
	ec.defSubr1(nil, "functionp", es.functionp_autogen, 1)
	ec.defSubrU(nil, "if", es.if_autogen, 2)
	ec.defSubr2(nil, "internal--define-uninitialized-variable", es.internalDefineUninitializedVariable_autogen, 1)
	ec.defSubr1(nil, "internal-make-var-non-special", es.makeVarNonSpecial_autogen, 1)
	ec.defSubrU(nil, "let", es.let_autogen, 1)
	ec.defSubrU(nil, "let*", es.letX_autogen, 1)
	ec.defSubr2(nil, "macroexpand", es.macroexpand_autogen, 1)
	ec.defSubr2(nil, "mapbacktrace", es.mapbacktrace_autogen, 1)
	ec.defSubrU(nil, "or", es.or_autogen, 0)
	ec.defSubrU(nil, "prog1", es.prog1_autogen, 1)
	ec.defSubrU(nil, "progn", es.progn_autogen, 0)
	ec.defSubrU(nil, "quote", es.quote_autogen, 1)
	ec.defSubrM(nil, "run-hook-with-args", es.runHookWithArgs_autogen, 1)
	ec.defSubrM(nil, "run-hook-with-args-until-failure", es.runHookWithArgsUntilFailure_autogen, 1)
	ec.defSubrM(nil, "run-hook-with-args-until-success", es.runHookWithArgsUntilSuccess_autogen, 1)
	ec.defSubrM(nil, "run-hook-wrapped", es.runHookWrapped_autogen, 2)
	ec.defSubrM(nil, "run-hooks", es.runHooks_autogen, 0)
	ec.defSubr2(nil, "set-default-toplevel-value", es.setDefaultToplevelValue_autogen, 2)
	ec.defSubrU(nil, "setq", es.setq_autogen, 0)
	ec.defSubr2(nil, "signal", es.signal_autogen, 2)
	ec.defSubr1(nil, "special-variable-p", es.specialVariableP_autogen, 1)
	ec.defSubr2(nil, "throw", es.throw_autogen, 2)
	ec.defSubrU(nil, "unwind-protect", es.unwindProtect_autogen, 1)
	ec.defSubrU(nil, "while", es.while_autogen, 1)
	ec.defSubr2(nil, "access-file", es.accessFile_autogen, 2)
	ec.defSubr3(nil, "add-name-to-file", es.addNameToFile_autogen, 2)
	ec.defSubr2(nil, "car-less-than-car", es.carLessThanCar_autogen, 2)
	ec.defSubr0(nil, "clear-buffer-auto-save-failure", es.clearBufferAutoSaveFailure_autogen)
	ec.defSubr6(nil, "copy-file", es.copyFile_autogen, 2)
	ec.defSubr0(nil, "default-file-modes", es.defaultFileModes_autogen)
	ec.defSubr1(nil, "delete-directory-internal", es.deleteDirectoryInternal_autogen, 1)
	ec.defSubr1(nil, "delete-file-internal", es.deleteFileInternal_autogen, 1)
	ec.defSubr1(nil, "directory-file-name", es.directoryFileName_autogen, 1)
	ec.defSubr1(nil, "directory-name-p", es.directoryNameP_autogen, 1)
	ec.defSubr2(nil, "do-auto-save", es.doAutoSave_autogen, 0)
	ec.defSubr2(nil, "expand-file-name", es.expandFileName_autogen, 1)
	ec.defSubr1(nil, "file-accessible-directory-p", es.fileAccessibleDirectoryP_autogen, 1)
	ec.defSubr1(nil, "file-acl", es.fileAcl_autogen, 1)
	ec.defSubr1(nil, "file-directory-p", es.fileDirectoryP_autogen, 1)
	ec.defSubr1(nil, "file-executable-p", es.fileExecutableP_autogen, 1)
	ec.defSubr1(nil, "file-exists-p", es.fileExistsP_autogen, 1)
	ec.defSubr2(nil, "file-modes", es.fileModes_autogen, 1)
	ec.defSubr1(nil, "file-name-absolute-p", es.fileNameAbsoluteP_autogen, 1)
	ec.defSubr1(nil, "file-name-as-directory", es.fileNameAsDirectory_autogen, 1)
	ec.defSubr1(nil, "file-name-case-insensitive-p", es.fileNameCaseInsensitiveP_autogen, 1)
	ec.defSubrM(nil, "file-name-concat", es.fileNameConcat_autogen, 1)
	ec.defSubr1(nil, "file-name-directory", es.fileNameDirectory_autogen, 1)
	ec.defSubr1(nil, "file-name-nondirectory", es.fileNameNondirectory_autogen, 1)
	ec.defSubr2(nil, "file-newer-than-file-p", es.fileNewerThanFileP_autogen, 2)
	ec.defSubr1(nil, "file-readable-p", es.fileReadableP_autogen, 1)
	ec.defSubr1(nil, "file-regular-p", es.fileRegularP_autogen, 1)
	ec.defSubr1(nil, "file-selinux-context", es.fileSelinuxContext_autogen, 1)
	ec.defSubr1(nil, "file-symlink-p", es.fileSymlinkP_autogen, 1)
	ec.defSubr1(nil, "file-system-info", es.fileSystemInfo_autogen, 1)
	ec.defSubr1(nil, "file-system-info", es.fileSystemInfo_1_autogen, 1)
	ec.defSubr1(nil, "file-writable-p", es.fileWritableP_autogen, 1)
	ec.defSubr2(nil, "find-file-name-handler", es.findFileNameHandler_autogen, 2)
	ec.defSubr5(nil, "insert-file-contents", es.insertFileContents_autogen, 1)
	ec.defSubr1(nil, "make-directory-internal", es.makeDirectoryInternal_autogen, 1)
	ec.defSubr3(nil, "make-symbolic-link", es.makeSymbolicLink_autogen, 2)
	ec.defSubr4(nil, "make-temp-file-internal", es.makeTempFileInternal_autogen, 4)
	ec.defSubr1(nil, "make-temp-name", es.makeTempName_autogen, 1)
	ec.defSubr0(nil, "next-read-file-uses-dialog-p", es.nextReadFileUsesDialogP_autogen)
	ec.defSubr0(nil, "recent-auto-save-p", es.recentAutoSaveP_autogen)
	ec.defSubr3(nil, "rename-file", es.renameFile_autogen, 2)
	ec.defSubr2(nil, "set-binary-mode", es.setBinaryMode_autogen, 2)
	ec.defSubr0(nil, "set-buffer-auto-saved", es.setBufferAutoSaved_autogen)
	ec.defSubr1(nil, "set-default-file-modes", es.setDefaultFileModes_autogen, 1)
	ec.defSubr2(nil, "set-file-acl", es.setFileAcl_autogen, 2)
	ec.defSubr3(nil, "set-file-modes", es.setFileModes_autogen, 2)
	ec.defSubr2(nil, "set-file-selinux-context", es.setFileSelinuxContext_autogen, 2)
	ec.defSubr3(nil, "set-file-times", es.setFileTimes_autogen, 1)
	ec.defSubr1(nil, "set-visited-file-modtime", es.setVisitedFileModtime_autogen, 0)
	ec.defSubr1(nil, "substitute-in-file-name", es.substituteInFileName_autogen, 1)
	ec.defSubr1(nil, "unhandled-file-name-directory", es.unhandledFileNameDirectory_autogen, 1)
	ec.defSubr0(nil, "unix-sync", es.unixSync_autogen)
	ec.defSubr1(nil, "verify-visited-file-modtime", es.verifyVisitedFileModtime_autogen, 0)
	ec.defSubr0(nil, "visited-file-modtime", es.visitedFileModtime_autogen)
	ec.defSubr7(nil, "write-region", es.writeRegion_autogen, 3)
	ec.defSubr1(nil, "file-locked-p", es.fileLockedP_autogen, 1)
	ec.defSubr1(nil, "lock-buffer", es.lockBuffer_autogen, 0)
	ec.defSubr1(nil, "lock-file", es.lockFile_autogen, 1)
	ec.defSubr0(nil, "unlock-buffer", es.unlockBuffer_autogen)
	ec.defSubr1(nil, "unlock-file", es.unlockFile_autogen, 1)
	ec.defSubr1(nil, "abs", es.abs_autogen, 1)
	ec.defSubr1(nil, "acos", es.acos_autogen, 1)
	ec.defSubr1(nil, "asin", es.asin_autogen, 1)
	ec.defSubr2(nil, "atan", es.atan_autogen, 1)
	ec.defSubr2(nil, "ceiling", es.ceiling_autogen, 1)
	ec.defSubr2(nil, "copysign", es.copysign_autogen, 2)
	ec.defSubr1(nil, "cos", es.cos_autogen, 1)
	ec.defSubr1(nil, "exp", es.exp_autogen, 1)
	ec.defSubr2(nil, "expt", es.expt_autogen, 2)
	ec.defSubr1(nil, "fceiling", es.fceiling_autogen, 1)
	ec.defSubr1(nil, "ffloor", es.ffloor_autogen, 1)
	ec.defSubr1(nil, "float", es.float_autogen, 1)
	ec.defSubr2(nil, "floor", es.floor_autogen, 1)
	ec.defSubr1(nil, "frexp", es.frexp_autogen, 1)
	ec.defSubr1(nil, "fround", es.fround_autogen, 1)
	ec.defSubr1(nil, "ftruncate", es.ftruncate_autogen, 1)
	ec.defSubr1(nil, "isnan", es.isnan_autogen, 1)
	ec.defSubr2(nil, "ldexp", es.ldexp_autogen, 2)
	ec.defSubr2(nil, "log", es.log_autogen, 1)
	ec.defSubr1(nil, "logb", es.logb_autogen, 1)
	ec.defSubr2(nil, "round", es.round_autogen, 1)
	ec.defSubr1(nil, "sin", es.sin_autogen, 1)
	ec.defSubr1(nil, "sqrt", es.sqrt_autogen, 1)
	ec.defSubr1(nil, "tan", es.tan_autogen, 1)
	ec.defSubr2(nil, "truncate", es.truncate_autogen, 1)
	ec.defSubrM(nil, "append", es.append_autogen, 0)
	ec.defSubr3(nil, "assoc", es.assoc_autogen, 2)
	ec.defSubr2(nil, "assq", es.assq_autogen, 2)
	ec.defSubr4(nil, "base64-decode-region", es.base64DecodeRegion_autogen, 2)
	ec.defSubr3(nil, "base64-decode-string", es.base64DecodeString_autogen, 1)
	ec.defSubr3(nil, "base64-encode-region", es.base64EncodeRegion_autogen, 2)
	ec.defSubr2(nil, "base64-encode-string", es.base64EncodeString_autogen, 1)
	ec.defSubr3(nil, "base64url-encode-region", es.base64urlEncodeRegion_autogen, 2)
	ec.defSubr2(nil, "base64url-encode-string", es.base64urlEncodeString_autogen, 1)
	ec.defSubr1(nil, "buffer-hash", es.bufferHash_autogen, 0)
	ec.defSubr1(nil, "buffer-line-statistics", es.bufferLineStatistics_autogen, 0)
	ec.defSubr1(nil, "clear-string", es.clearString_autogen, 1)
	ec.defSubr1(nil, "clrhash", es.clrhash_autogen, 1)
	ec.defSubr7(nil, "compare-strings", es.compareStrings_autogen, 6)
	ec.defSubrM(nil, "concat", es.concat_autogen, 0)
	ec.defSubr1(nil, "copy-alist", es.copyAlist_autogen, 1)
	ec.defSubr1(nil, "copy-hash-table", es.copyHashTable_autogen, 1)
	ec.defSubr1(nil, "copy-sequence", es.copySequence_autogen, 1)
	ec.defSubr3(nil, "define-hash-table-test", es.defineHashTableTest_autogen, 3)
	ec.defSubr2(nil, "delete", es.delete_autogen, 2)
	ec.defSubr2(nil, "delq", es.delq_autogen, 2)
	ec.defSubr2(nil, "elt", es.elt_autogen, 2)
	ec.defSubr2(nil, "eql", es.eql_autogen, 2)
	ec.defSubr2(nil, "equal", es.equal_autogen, 2)
	ec.defSubr2(nil, "equal-including-properties", es.equalIncludingProperties_autogen, 2)
	ec.defSubr2(nil, "featurep", es.featurep_autogen, 1)
	ec.defSubr2(nil, "fillarray", es.fillarray_autogen, 2)
	ec.defSubr2(nil, "get", es.get_autogen, 2)
	ec.defSubr3(nil, "gethash", es.gethash_autogen, 2)
	ec.defSubr1(nil, "hash-table-count", es.hashTableCount_autogen, 1)
	ec.defSubr1(nil, "hash-table-p", es.hashTableP_autogen, 1)
	ec.defSubr1(nil, "hash-table-rehash-size", es.hashTableRehashSize_autogen, 1)
	ec.defSubr1(nil, "hash-table-rehash-threshold", es.hashTableRehashThreshold_autogen, 1)
	ec.defSubr1(nil, "hash-table-size", es.hashTableSize_autogen, 1)
	ec.defSubr1(nil, "hash-table-test", es.hashTableTest_autogen, 1)
	ec.defSubr1(nil, "hash-table-weakness", es.hashTableWeakness_autogen, 1)
	ec.defSubr1(nil, "identity", es.identity_autogen, 1)
	ec.defSubr1(nil, "length", es.length_autogen, 1)
	ec.defSubr2(nil, "length<", es.lengthLess_autogen, 2)
	ec.defSubr2(nil, "length=", es.lengthEqual_autogen, 2)
	ec.defSubr2(nil, "length>", es.lengthGreater_autogen, 2)
	ec.defSubr2(nil, "line-number-at-pos", es.lineNumberAtPos_autogen, 0)
	ec.defSubr1(nil, "load-average", es.loadAverage_autogen, 0)
	ec.defSubr1(nil, "locale-info", es.localeInfo_autogen, 1)
	ec.defSubrM(nil, "make-hash-table", es.makeHashTable_autogen, 0)
	ec.defSubr2(nil, "mapc", es.mapc_autogen, 2)
	ec.defSubr2(nil, "mapcan", es.mapcan_autogen, 2)
	ec.defSubr2(nil, "mapcar", es.mapcar_autogen, 2)
	ec.defSubr3(nil, "mapconcat", es.mapconcat_autogen, 2)
	ec.defSubr2(nil, "maphash", es.maphash_autogen, 2)
	ec.defSubr5(nil, "md5", es.md5_autogen, 1)
	ec.defSubr2(nil, "member", es.member_autogen, 2)
	ec.defSubr2(nil, "memq", es.memq_autogen, 2)
	ec.defSubr2(nil, "memql", es.memql_autogen, 2)
	ec.defSubrM(nil, "nconc", es.nconc_autogen, 0)
	ec.defSubr1(nil, "nreverse", es.nreverse_autogen, 1)
	ec.defSubr2(nil, "ntake", es.ntake_autogen, 2)
	ec.defSubr2(nil, "nth", es.nth_autogen, 2)
	ec.defSubr2(nil, "nthcdr", es.nthcdr_autogen, 2)
	ec.defSubr1(nil, "object-intervals", es.objectIntervals_autogen, 1)
	ec.defSubr3(nil, "plist-get", es.plistGet_autogen, 2)
	ec.defSubr3(nil, "plist-member", es.plistMember_autogen, 2)
	ec.defSubr4(nil, "plist-put", es.plistPut_autogen, 3)
	ec.defSubr1(nil, "proper-list-p", es.properListP_autogen, 1)
	ec.defSubr2(nil, "provide", es.provide_autogen, 1)
	ec.defSubr3(nil, "put", es.put_autogen, 3)
	ec.defSubr3(nil, "puthash", es.puthash_autogen, 3)
	ec.defSubr1(nil, "random", es.random_autogen, 0)
	ec.defSubr2(nil, "rassoc", es.rassoc_autogen, 2)
	ec.defSubr2(nil, "rassq", es.rassq_autogen, 2)
	ec.defSubr2(nil, "remhash", es.remhash_autogen, 2)
	ec.defSubr3(nil, "require", es.require_autogen, 1)
	ec.defSubr1(nil, "reverse", es.reverse_autogen, 1)
	ec.defSubr1(nil, "safe-length", es.safeLength_autogen, 1)
	ec.defSubr5(nil, "secure-hash", es.secureHash_autogen, 2)
	ec.defSubr0(nil, "secure-hash-algorithms", es.secureHashAlgorithms_autogen)
	ec.defSubr2(nil, "sort", es.sort_autogen, 2)
	ec.defSubr1(nil, "string-as-multibyte", es.stringAsMultibyte_autogen, 1)
	ec.defSubr1(nil, "string-as-unibyte", es.stringAsUnibyte_autogen, 1)
	ec.defSubr1(nil, "string-bytes", es.stringBytes_autogen, 1)
	ec.defSubr4(nil, "string-collate-equalp", es.stringCollateEqualp_autogen, 2)
	ec.defSubr4(nil, "string-collate-lessp", es.stringCollateLessp_autogen, 2)
	ec.defSubr3(nil, "string-distance", es.stringDistance_autogen, 2)
	ec.defSubr2(nil, "string-equal", es.stringEqual_autogen, 2)
	ec.defSubr2(nil, "string-lessp", es.stringLessp_autogen, 2)
	ec.defSubr1(nil, "string-make-multibyte", es.stringMakeMultibyte_autogen, 1)
	ec.defSubr1(nil, "string-make-unibyte", es.stringMakeUnibyte_autogen, 1)
	ec.defSubr3(nil, "string-search", es.stringSearch_autogen, 2)
	ec.defSubr1(nil, "string-to-multibyte", es.stringToMultibyte_autogen, 1)
	ec.defSubr1(nil, "string-to-unibyte", es.stringToUnibyte_autogen, 1)
	ec.defSubr2(nil, "string-version-lessp", es.stringVersionLessp_autogen, 2)
	ec.defSubr3(nil, "substring", es.substring_autogen, 1)
	ec.defSubr3(nil, "substring-no-properties", es.substringNoProperties_autogen, 1)
	ec.defSubr1(nil, "sxhash-eq", es.sxhashEq_autogen, 1)
	ec.defSubr1(nil, "sxhash-eql", es.sxhashEql_autogen, 1)
	ec.defSubr1(nil, "sxhash-equal", es.sxhashEqual_autogen, 1)
	ec.defSubr1(nil, "sxhash-equal-including-properties", es.sxhashEqualIncludingProperties_autogen, 1)
	ec.defSubr2(nil, "take", es.take_autogen, 2)
	ec.defSubrM(nil, "vconcat", es.vconcat_autogen, 0)
	ec.defSubrM(nil, "widget-apply", es.widgetApply_autogen, 2)
	ec.defSubr2(nil, "widget-get", es.widgetGet_autogen, 2)
	ec.defSubr3(nil, "widget-put", es.widgetPut_autogen, 3)
	ec.defSubr1(nil, "yes-or-no-p", es.yesOrNoP_autogen, 1)
	ec.defSubr0(nil, "clear-font-cache", es.clearFontCache_autogen)
	ec.defSubr2(nil, "close-font", es.closeFont_autogen, 1)
	ec.defSubr2(nil, "draw-string", es.drawString_autogen, 2)
	ec.defSubr2(nil, "find-font", es.findFont_autogen, 1)
	ec.defSubr3(nil, "font-at", es.fontAt_autogen, 1)
	ec.defSubr6(nil, "font-drive-otf", es.fontDriveOtf_autogen, 6)
	ec.defSubr2(nil, "font-face-attributes", es.fontFaceAttributes_autogen, 1)
	ec.defSubr1(nil, "font-family-list", es.fontFamilyList_autogen, 0)
	ec.defSubr2(nil, "font-get", es.fontGet_autogen, 2)
	ec.defSubr4(nil, "font-get-glyphs", es.fontGetGlyphs_autogen, 3)
	ec.defSubr3(nil, "font-has-char-p", es.fontHasCharP_autogen, 2)
	ec.defSubr2(nil, "font-info", es.fontInfo_autogen, 1)
	ec.defSubr2(nil, "font-match-p", es.fontMatchP_autogen, 2)
	ec.defSubr3(nil, "font-otf-alternates", es.fontOtfAlternates_autogen, 3)
	ec.defSubr3(nil, "font-put", es.fontPut_autogen, 3)
	ec.defSubr2(nil, "font-shape-gstring", es.fontShapeGstring_autogen, 2)
	ec.defSubrM(nil, "font-spec", es.fontSpec_autogen, 0)
	ec.defSubr2(nil, "font-variation-glyphs", es.fontVariationGlyphs_autogen, 2)
	ec.defSubr2(nil, "font-xlfd-name", es.fontXlfdName_autogen, 1)
	ec.defSubr2(nil, "fontp", es.fontp_autogen, 1)
	ec.defSubr1(nil, "frame-font-cache", es.frameFontCache_autogen, 0)
	ec.defSubr2(nil, "internal-char-font", es.internalCharFont_autogen, 1)
	ec.defSubr4(nil, "list-fonts", es.listFonts_autogen, 1)
	ec.defSubr3(nil, "open-font", es.openFont_autogen, 1)
	ec.defSubr1(nil, "query-font", es.queryFont_autogen, 1)
	ec.defSubr3(nil, "fontset-font", es.fontsetFont_autogen, 2)
	ec.defSubr2(nil, "fontset-info", es.fontsetInfo_autogen, 1)
	ec.defSubr0(nil, "fontset-list", es.fontsetList_autogen)
	ec.defSubr0(nil, "fontset-list-all", es.fontsetListAll_autogen)
	ec.defSubr2(nil, "new-fontset", es.newFontset_autogen, 2)
	ec.defSubr2(nil, "query-fontset", es.queryFontset_autogen, 1)
	ec.defSubr5(nil, "set-fontset-font", es.setFontsetFont_autogen, 3)
	ec.defSubr2(nil, "delete-frame", es.deleteFrame_autogen, 0)
	ec.defSubr2(nil, "frame--set-was-invisible", es.frameSetWasInvisible_autogen, 2)
	ec.defSubr2(nil, "frame-after-make-frame", es.frameAfterMakeFrame_autogen, 2)
	ec.defSubr2(nil, "frame-ancestor-p", es.frameAncestorP_autogen, 2)
	ec.defSubr1(nil, "frame-bottom-divider-width", es.bottomDividerWidth_autogen, 0)
	ec.defSubr1(nil, "frame-char-height", es.frameCharHeight_autogen, 0)
	ec.defSubr1(nil, "frame-char-width", es.frameCharWidth_autogen, 0)
	ec.defSubr1(nil, "frame-child-frame-border-width", es.frameChildFrameBorderWidth_autogen, 0)
	ec.defSubr1(nil, "frame-focus", es.frameFocus_autogen, 0)
	ec.defSubr1(nil, "frame-fringe-width", es.fringeWidth_autogen, 0)
	ec.defSubr1(nil, "frame-internal-border-width", es.frameInternalBorderWidth_autogen, 0)
	ec.defSubr0(nil, "frame-list", es.frameList_autogen)
	ec.defSubr1(nil, "frame-live-p", es.frameLiveP_autogen, 1)
	ec.defSubr1(nil, "frame-native-height", es.frameNativeHeight_autogen, 0)
	ec.defSubr1(nil, "frame-native-width", es.frameNativeWidth_autogen, 0)
	ec.defSubr2(nil, "frame-parameter", es.frameParameter_autogen, 2)
	ec.defSubr1(nil, "frame-parameters", es.frameParameters_autogen, 0)
	ec.defSubr1(nil, "frame-parent", es.frameParent_autogen, 0)
	ec.defSubr1(nil, "frame-pointer-visible-p", es.framePointerVisibleP_autogen, 0)
	ec.defSubr1(nil, "frame-position", es.framePosition_autogen, 0)
	ec.defSubr1(nil, "frame-right-divider-width", es.rightDividerWidth_autogen, 0)
	ec.defSubr1(nil, "frame-scale-factor", es.frameScaleFactor_autogen, 0)
	ec.defSubr1(nil, "frame-scroll-bar-height", es.scrollBarHeight_autogen, 0)
	ec.defSubr1(nil, "frame-scroll-bar-width", es.scrollBarWidth_autogen, 0)
	ec.defSubr1(nil, "frame-text-cols", es.frameTextCols_autogen, 0)
	ec.defSubr1(nil, "frame-text-height", es.frameTextHeight_autogen, 0)
	ec.defSubr1(nil, "frame-text-lines", es.frameTextLines_autogen, 0)
	ec.defSubr1(nil, "frame-text-width", es.frameTextWidth_autogen, 0)
	ec.defSubr1(nil, "frame-total-cols", es.frameTotalCols_autogen, 0)
	ec.defSubr1(nil, "frame-total-lines", es.frameTotalLines_autogen, 0)
	ec.defSubr1(nil, "frame-visible-p", es.frameVisibleP_autogen, 1)
	ec.defSubr1(nil, "frame-window-state-change", es.frameWindowStateChange_autogen, 0)
	ec.defSubr4(nil, "frame-windows-min-size", es.frameWindowsMinSize_autogen, 4)
	ec.defSubr1(nil, "framep", es.framep_autogen, 1)
	ec.defSubr1(nil, "handle-switch-frame", es.handleSwitchFrame_autogen, 1)
	ec.defSubr1(nil, "iconify-frame", es.iconifyFrame_autogen, 0)
	ec.defSubr0(nil, "last-nonminibuffer-frame", es.lastNonminibufFrame_autogen)
	ec.defSubr1(nil, "lower-frame", es.lowerFrame_autogen, 0)
	ec.defSubr2(nil, "make-frame-invisible", es.makeFrameInvisible_autogen, 0)
	ec.defSubr1(nil, "make-frame-visible", es.makeFrameVisible_autogen, 0)
	ec.defSubr1(nil, "make-terminal-frame", es.makeTerminalFrame_autogen, 1)
	ec.defSubr2(nil, "modify-frame-parameters", es.modifyFrameParameters_autogen, 2)
	ec.defSubr0(nil, "mouse-pixel-position", es.mousePixelPosition_autogen)
	ec.defSubr0(nil, "mouse-position", es.mousePosition_autogen)
	ec.defSubr2(nil, "next-frame", es.nextFrame_autogen, 0)
	ec.defSubr0(nil, "old-selected-frame", es.oldSelectedFrame_autogen)
	ec.defSubr2(nil, "previous-frame", es.previousFrame_autogen, 0)
	ec.defSubr1(nil, "raise-frame", es.raiseFrame_autogen, 0)
	ec.defSubr1(nil, "reconsider-frame-fonts", es.reconsiderFrameFonts_autogen, 1)
	ec.defSubr2(nil, "redirect-frame-focus", es.redirectFrameFocus_autogen, 1)
	ec.defSubr2(nil, "select-frame", es.selectFrame_autogen, 1)
	ec.defSubr0(nil, "selected-frame", es.selectedFrame_autogen)
	ec.defSubr4(nil, "set-frame-height", es.setFrameHeight_autogen, 2)
	ec.defSubr3(nil, "set-frame-position", es.setFramePosition_autogen, 3)
	ec.defSubr4(nil, "set-frame-size", es.setFrameSize_autogen, 3)
	ec.defSubr4(nil, "set-frame-width", es.setFrameWidth_autogen, 2)
	ec.defSubr2(nil, "set-frame-window-state-change", es.setFrameWindowStateChange_autogen, 0)
	ec.defSubr3(nil, "set-mouse-pixel-position", es.setMousePixelPosition_autogen, 3)
	ec.defSubr3(nil, "set-mouse-position", es.setMousePosition_autogen, 3)
	ec.defSubr1(nil, "tool-bar-pixel-width", es.toolBarPixelWidth_autogen, 0)
	ec.defSubr0(nil, "visible-frame-list", es.visibleFrameList_autogen)
	ec.defSubr1(nil, "window-system", es.windowSystem_autogen, 0)
	ec.defSubr2(nil, "x-focus-frame", es.xFocusFrame_autogen, 1)
	ec.defSubr4(nil, "x-get-resource", es.xGetResource_autogen, 2)
	ec.defSubr1(nil, "x-parse-geometry", es.xParseGeometry_autogen, 1)
	ec.defSubr5(nil, "define-fringe-bitmap", es.defineFringeBitmap_autogen, 2)
	ec.defSubr1(nil, "destroy-fringe-bitmap", es.destroyFringeBitmap_autogen, 1)
	ec.defSubr2(nil, "fringe-bitmaps-at-pos", es.fringeBitmapsAtPos_autogen, 0)
	ec.defSubr2(nil, "set-fringe-bitmap-face", es.setFringeBitmapFace_autogen, 1)
	ec.defSubr3(nil, "gfile-add-watch", es.gfileAddWatch_autogen, 3)
	ec.defSubr1(nil, "gfile-monitor-name", es.gfileMonitorName_autogen, 1)
	ec.defSubr1(nil, "gfile-rm-watch", es.gfileRmWatch_autogen, 1)
	ec.defSubr1(nil, "gfile-valid-p", es.gfileValidP_autogen, 1)
	ec.defSubr2(nil, "gnutls-asynchronous-parameters", es.gnutlsAsynchronousParameters_autogen, 2)
	ec.defSubr0(nil, "gnutls-available-p", es.gnutlsAvailableP_autogen)
	ec.defSubr3(nil, "gnutls-boot", es.gnutlsBoot_autogen, 3)
	ec.defSubr2(nil, "gnutls-bye", es.gnutlsBye_autogen, 2)
	ec.defSubr0(nil, "gnutls-ciphers", es.gnutlsCiphers_autogen)
	ec.defSubr1(nil, "gnutls-deinit", es.gnutlsDeinit_autogen, 1)
	ec.defSubr0(nil, "gnutls-digests", es.gnutlsDigests_autogen)
	ec.defSubr1(nil, "gnutls-error-fatalp", es.gnutlsErrorFatalp_autogen, 1)
	ec.defSubr1(nil, "gnutls-error-string", es.gnutlsErrorString_autogen, 1)
	ec.defSubr1(nil, "gnutls-errorp", es.gnutlsErrorp_autogen, 1)
	ec.defSubr1(nil, "gnutls-format-certificate", es.gnutlsFormatCertificate_autogen, 1)
	ec.defSubr1(nil, "gnutls-get-initstage", es.gnutlsGetInitstage_autogen, 1)
	ec.defSubr2(nil, "gnutls-hash-digest", es.gnutlsHashDigest_autogen, 2)
	ec.defSubr3(nil, "gnutls-hash-mac", es.gnutlsHashMac_autogen, 3)
	ec.defSubr0(nil, "gnutls-macs", es.gnutlsMacs_autogen)
	ec.defSubr1(nil, "gnutls-peer-status", es.gnutlsPeerStatus_autogen, 1)
	ec.defSubr1(nil, "gnutls-peer-status-warning-describe", es.gnutlsPeerStatusWarningDescribe_autogen, 1)
	ec.defSubr5(nil, "gnutls-symmetric-decrypt", es.gnutlsSymmetricDecrypt_autogen, 4)
	ec.defSubr5(nil, "gnutls-symmetric-encrypt", es.gnutlsSymmetricEncrypt_autogen, 4)
	ec.defSubr2(nil, "clear-image-cache", es.clearImageCache_autogen, 0)
	ec.defSubr0(nil, "image-cache-size", es.imageCacheSize_autogen)
	ec.defSubr2(nil, "image-flush", es.imageFlush_autogen, 1)
	ec.defSubr2(nil, "image-mask-p", es.imageMaskP_autogen, 1)
	ec.defSubr2(nil, "image-metadata", es.imageMetadata_autogen, 1)
	ec.defSubr3(nil, "image-size", es.imageSize_autogen, 1)
	ec.defSubr1(nil, "image-transforms-p", es.imageTransformsP_autogen, 0)
	ec.defSubr0(nil, "imagemagick-types", es.imagemagickTypes_autogen)
	ec.defSubr1(nil, "imagep", es.imagep_autogen, 1)
	ec.defSubr1(nil, "init-image-library", es.initImageLibrary_autogen, 1)
	ec.defSubr1(nil, "lookup-image", es.lookupImage_autogen, 1)
	ec.defSubr7(nil, "compute-motion", es.computeMotion_autogen, 7)
	ec.defSubr0(nil, "current-column", es.currentColumn_autogen)
	ec.defSubr0(nil, "current-indentation", es.currentIndentation_autogen)
	ec.defSubr2(nil, "indent-to", es.indentTo_autogen, 1)
	ec.defSubr1(nil, "line-number-display-width", es.lineNumberDisplayWidth_autogen, 0)
	ec.defSubr2(nil, "move-to-column", es.moveToColumn_autogen, 1)
	ec.defSubr3(nil, "vertical-motion", es.verticalMotion_autogen, 1)
	ec.defSubr3(nil, "inotify-add-watch", es.inotifyAddWatch_autogen, 3)
	ec.defSubr0(nil, "inotify-allocated-p", es.inotifyAllocatedP_autogen)
	ec.defSubr1(nil, "inotify-rm-watch", es.inotifyRmWatch_autogen, 1)
	ec.defSubr1(nil, "inotify-valid-p", es.inotifyValidP_autogen, 1)
	ec.defSubr0(nil, "inotify-watch-list", es.inotifyWatchList_autogen)
	ec.defSubr0(nil, "combine-after-change-execute", es.combineAfterChangeExecute_autogen)
	ec.defSubr0(nil, "json--available-p", es.jsonAvailableP_autogen)
	ec.defSubrM(nil, "json-insert", es.jsonInsert_autogen, 1)
	ec.defSubrM(nil, "json-parse-buffer", es.jsonParseBuffer_autogen, 0)
	ec.defSubrM(nil, "json-parse-string", es.jsonParseString_autogen, 1)
	ec.defSubrM(nil, "json-serialize", es.jsonSerialize_autogen, 1)
	ec.defSubr0(nil, "abort-recursive-edit", es.abortRecursiveEdit_autogen)
	ec.defSubr1(nil, "clear-this-command-keys", es.clearThisCommandKeys_autogen, 0)
	ec.defSubr3(nil, "command-error-default-function", es.commandErrorDefaultFunction_autogen, 3)
	ec.defSubr0(nil, "current-idle-time", es.currentIdleTime_autogen)
	ec.defSubr0(nil, "current-input-mode", es.currentInputMode_autogen)
	ec.defSubr0(nil, "discard-input", es.discardInput_autogen)
	ec.defSubr1(nil, "event-convert-list", es.eventConvertList_autogen, 1)
	ec.defSubr0(nil, "exit-recursive-edit", es.exitRecursiveEdit_autogen)
	ec.defSubr1(nil, "input-pending-p", es.inputPendingP_autogen, 0)
	ec.defSubr1(nil, "internal--track-mouse", es.internalTrackMouse_autogen, 1)
	ec.defSubr1(nil, "internal-event-symbol-parse-modifiers", es.eventSymbolParseModifiers_autogen, 1)
	ec.defSubr1(nil, "internal-handle-focus-in", es.internalHandleFocusIn_autogen, 1)
	ec.defSubr1(nil, "lossage-size", es.lossageSize_autogen, 0)
	ec.defSubr1(nil, "open-dribble-file", es.openDribbleFile_autogen, 1)
	ec.defSubr2(nil, "posn-at-point", es.posnAtPoint_autogen, 0)
	ec.defSubr4(nil, "posn-at-x-y", es.posnAtXY_autogen, 2)
	ec.defSubr6(nil, "read-key-sequence", es.readKeySequence_autogen, 1)
	ec.defSubr6(nil, "read-key-sequence-vector", es.readKeySequenceVector_autogen, 1)
	ec.defSubr1(nil, "recent-keys", es.recentKeys_autogen, 0)
	ec.defSubr0(nil, "recursion-depth", es.recursionDepth_autogen)
	ec.defSubr0(nil, "recursive-edit", es.recursiveEdit_autogen)
	ec.defSubr1(nil, "set--this-command-keys", es.setThisCommandKeys_autogen, 1)
	ec.defSubr1(nil, "set-input-interrupt-mode", es.setInputInterruptMode_autogen, 1)
	ec.defSubr2(nil, "set-input-meta-mode", es.setInputMetaMode_autogen, 1)
	ec.defSubr4(nil, "set-input-mode", es.setInputMode_autogen, 3)
	ec.defSubr2(nil, "set-output-flow-control", es.setOutputFlowControl_autogen, 1)
	ec.defSubr1(nil, "set-quit-char", es.setQuitChar_autogen, 1)
	ec.defSubr1(nil, "suspend-emacs", es.suspendEmacs_autogen, 0)
	ec.defSubr0(nil, "this-command-keys", es.thisCommandKeys_autogen)
	ec.defSubr0(nil, "this-command-keys-vector", es.thisCommandKeysVector_autogen)
	ec.defSubr0(nil, "this-single-command-keys", es.thisSingleCommandKeys_autogen)
	ec.defSubr0(nil, "this-single-command-raw-keys", es.thisSingleCommandRawKeys_autogen)
	ec.defSubr0(nil, "top-level", es.topLevel_autogen)
	ec.defSubr2(nil, "accessible-keymaps", es.accessibleKeymaps_autogen, 1)
	ec.defSubr3(nil, "command-remapping", es.commandRemapping_autogen, 1)
	ec.defSubr1(nil, "copy-keymap", es.copyKeymap_autogen, 1)
	ec.defSubr2(nil, "current-active-maps", es.currentActiveMaps_autogen, 0)
	ec.defSubr0(nil, "current-global-map", es.currentGlobalMap_autogen)
	ec.defSubr0(nil, "current-local-map", es.currentLocalMap_autogen)
	ec.defSubr0(nil, "current-minor-mode-maps", es.currentMinorModeMaps_autogen)
	ec.defSubr4(nil, "define-key", es.defineKey_autogen, 3)
	ec.defSubr3(nil, "describe-buffer-bindings", es.describeBufferBindings_autogen, 1)
	ec.defSubr2(nil, "describe-vector", es.describeVector_autogen, 1)
	ec.defSubr7(nil, "help--describe-vector", es.helpDescribeVector_autogen, 7)
	ec.defSubr4(nil, "key-binding", es.keyBinding_autogen, 1)
	ec.defSubr2(nil, "key-description", es.keyDescription_autogen, 1)
	ec.defSubr2(nil, "keymap--get-keyelt", es.keymapGetKeyelt_autogen, 2)
	ec.defSubr1(nil, "keymap-parent", es.keymapParent_autogen, 1)
	ec.defSubr1(nil, "keymap-prompt", es.keymapPrompt_autogen, 1)
	ec.defSubr1(nil, "keymapp", es.keymapp_autogen, 1)
	ec.defSubr3(nil, "lookup-key", es.lookupKey_autogen, 2)
	ec.defSubr1(nil, "make-keymap", es.makeKeymap_autogen, 0)
	ec.defSubr1(nil, "make-sparse-keymap", es.makeSparseKeymap_autogen, 0)
	ec.defSubr3(nil, "map-keymap", es.mapKeymap_autogen, 2)
	ec.defSubr2(nil, "map-keymap-internal", es.mapKeymapInternal_autogen, 2)
	ec.defSubr2(nil, "minor-mode-key-binding", es.minorModeKeyBinding_autogen, 1)
	ec.defSubr2(nil, "set-keymap-parent", es.setKeymapParent_autogen, 2)
	ec.defSubr2(nil, "single-key-description", es.singleKeyDescription_autogen, 1)
	ec.defSubr1(nil, "text-char-description", es.textCharDescription_autogen, 1)
	ec.defSubr1(nil, "use-global-map", es.useGlobalMap_autogen, 1)
	ec.defSubr1(nil, "use-local-map", es.useLocalMap_autogen, 1)
	ec.defSubr5(nil, "where-is-internal", es.whereIsInternal_autogen, 1)
	ec.defSubr3(nil, "kqueue-add-watch", es.kqueueAddWatch_autogen, 3)
	ec.defSubr1(nil, "kqueue-rm-watch", es.kqueueRmWatch_autogen, 1)
	ec.defSubr1(nil, "kqueue-valid-p", es.kqueueValidP_autogen, 1)
	ec.defSubr4(nil, "lcms-cam02-ucs", es.lcmsCam02Ucs_autogen, 2)
	ec.defSubr5(nil, "lcms-cie-de2000", es.lcmsCieDe2000_autogen, 2)
	ec.defSubr3(nil, "lcms-jab->jch", es.lcmsJabToJch_autogen, 1)
	ec.defSubr3(nil, "lcms-jch->jab", es.lcmsJchToJab_autogen, 1)
	ec.defSubr3(nil, "lcms-jch->xyz", es.lcmsJchToXyz_autogen, 1)
	ec.defSubr1(nil, "lcms-temp->white-point", es.lcmsTempToWhitePoint_autogen, 1)
	ec.defSubr3(nil, "lcms-xyz->jch", es.lcmsXyzToJch_autogen, 1)
	ec.defSubr0(nil, "lcms2-available-p", es.lcms2AvailableP_autogen)
	ec.defSubr5(nil, "eval-buffer", es.evalBuffer_autogen, 0)
	ec.defSubr4(nil, "eval-region", es.evalRegion_autogen, 2)
	ec.defSubr0(nil, "get-file-char", es.getFileChar_autogen)
	ec.defSubr0(nil, "get-load-suffixes", es.getLoadSuffixes_autogen)
	ec.defSubr2(nil, "intern", es.intern_autogen, 1)
	ec.defSubr2(nil, "intern-soft", es.internSoft_autogen, 1)
	ec.defSubr5(nil, "load", es.load_autogen, 1)
	ec.defSubr4(nil, "locate-file-internal", es.locateFileInternal_autogen, 2)
	ec.defSubr3(nil, "lread--substitute-object-in-subtree", es.lreadSubstituteObjectInSubtree_autogen, 3)
	ec.defSubr2(nil, "mapatoms", es.mapatoms_autogen, 1)
	ec.defSubr1(nil, "read", es.read_autogen, 0)
	ec.defSubr3(nil, "read-char", es.readChar_autogen, 0)
	ec.defSubr3(nil, "read-char-exclusive", es.readCharExclusive_autogen, 0)
	ec.defSubr3(nil, "read-event", es.readEvent_autogen, 0)
	ec.defSubr3(nil, "read-from-string", es.readFromString_autogen, 1)
	ec.defSubr1(nil, "read-positioning-symbols", es.readPositioningSymbols_autogen, 0)
	ec.defSubr2(nil, "unintern", es.unintern_autogen, 1)
	ec.defSubr2(nil, "call-last-kbd-macro", es.callLastKbdMacro_autogen, 0)
	ec.defSubr0(nil, "cancel-kbd-macro-events", es.cancelKbdMacroEvents_autogen)
	ec.defSubr2(nil, "end-kbd-macro", es.endKbdMacro_autogen, 0)
	ec.defSubr3(nil, "execute-kbd-macro", es.executeKbdMacro_autogen, 1)
	ec.defSubr2(nil, "start-kbd-macro", es.startKbdMacro_autogen, 1)
	ec.defSubr1(nil, "store-kbd-macro-event", es.storeKbdMacroEvent_autogen, 1)
	ec.defSubr2(nil, "copy-marker", es.copyMarker_autogen, 0)
	ec.defSubr1(nil, "marker-buffer", es.markerBuffer_autogen, 1)
	ec.defSubr1(nil, "marker-insertion-type", es.markerInsertionType_autogen, 1)
	ec.defSubr1(nil, "marker-position", es.markerPosition_autogen, 1)
	ec.defSubr3(nil, "set-marker", es.setMarker_autogen, 2)
	ec.defSubr2(nil, "set-marker-insertion-type", es.setMarkerInsertionType_autogen, 2)
	ec.defSubr3(nil, "menu-bar-menu-at-x-y", es.menuBarMenuAtXY_autogen, 2)
	ec.defSubr3(nil, "x-popup-dialog", es.xPopupDialog_autogen, 2)
	ec.defSubr2(nil, "x-popup-menu", es.xPopupMenu_autogen, 2)
	ec.defSubr0(nil, "abort-minibuffers", es.abortMinibuffers_autogen)
	ec.defSubr0(nil, "active-minibuffer-window", es.activeMinibufferWindow_autogen)
	ec.defSubr4(nil, "all-completions", es.allCompletions_autogen, 2)
	ec.defSubr3(nil, "assoc-string", es.assocString_autogen, 2)
	ec.defSubr8(nil, "completing-read", es.completingRead_autogen, 2)
	ec.defSubr1(nil, "innermost-minibuffer-p", es.innermostMinibufferP_autogen, 0)
	ec.defSubr3(nil, "internal-complete-buffer", es.internalCompleteBuffer_autogen, 3)
	ec.defSubr0(nil, "minibuffer-contents", es.minibufferContents_autogen)
	ec.defSubr0(nil, "minibuffer-contents-no-properties", es.minibufferContentsNoProperties_autogen)
	ec.defSubr0(nil, "minibuffer-depth", es.minibufferDepth_autogen)
	ec.defSubr1(nil, "minibuffer-innermost-command-loop-p", es.minibufferInnermostCommandLoopP_autogen, 0)
	ec.defSubr0(nil, "minibuffer-prompt", es.minibufferPrompt_autogen)
	ec.defSubr0(nil, "minibuffer-prompt-end", es.minibufferPromptEnd_autogen)
	ec.defSubr2(nil, "minibufferp", es.minibufferp_autogen, 0)
	ec.defSubr4(nil, "read-buffer", es.readBuffer_autogen, 1)
	ec.defSubr2(nil, "read-command", es.readCommand_autogen, 1)
	ec.defSubr7(nil, "read-from-minibuffer", es.readFromMinibuffer_autogen, 1)
	ec.defSubr1(nil, "read-function", es.readFunction_autogen, 1)
	ec.defSubr5(nil, "read-string", es.readString_autogen, 1)
	ec.defSubr2(nil, "read-variable", es.readVariable_autogen, 1)
	ec.defSubr1(nil, "set-minibuffer-window", es.setMinibufferWindow_autogen, 1)
	ec.defSubr3(nil, "test-completion", es.testCompletion_autogen, 2)
	ec.defSubr3(nil, "try-completion", es.tryCompletion_autogen, 2)
	ec.defSubr2(nil, "dump-emacs-portable", es.dumpEmacsPortable_autogen, 1)
	ec.defSubr2(nil, "dump-emacs-portable--sort-predicate", es.dumpEmacsPortableSortPredicate_autogen, 2)
	ec.defSubr2(nil, "dump-emacs-portable--sort-predicate-copied", es.dumpEmacsPortableSortPredicateCopied_autogen, 2)
	ec.defSubr0(nil, "pdumper-stats", es.pdumperStats_autogen)
	ec.defSubr1(nil, "pgtk-backend-display-class", es.pgtkBackendDisplayClass_autogen, 0)
	ec.defSubr1(nil, "pgtk-display-monitor-attributes-list", es.pgtkDisplayMonitorAttributesList_autogen, 0)
	ec.defSubr1(nil, "pgtk-font-name", es.pgtkFontName_autogen, 1)
	ec.defSubr2(nil, "pgtk-frame-edges", es.pgtkFrameEdges_autogen, 0)
	ec.defSubr1(nil, "pgtk-frame-geometry", es.pgtkFrameGeometry_autogen, 0)
	ec.defSubr3(nil, "pgtk-frame-restack", es.pgtkFrameRestack_autogen, 2)
	ec.defSubr0(nil, "pgtk-get-page-setup", es.pgtkGetPageSetup_autogen)
	ec.defSubr0(nil, "pgtk-mouse-absolute-pixel-position", es.pgtkMouseAbsolutePixelPosition_autogen)
	ec.defSubr0(nil, "pgtk-page-setup-dialog", es.pgtkPageSetupDialog_autogen)
	ec.defSubr1(nil, "pgtk-print-frames-dialog", es.pgtkPrintFramesDialog_autogen, 0)
	ec.defSubr2(nil, "pgtk-set-monitor-scale-factor", es.pgtkSetMonitorScaleFactor_autogen, 2)
	ec.defSubr2(nil, "pgtk-set-mouse-absolute-pixel-position", es.pgtkSetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr2(nil, "pgtk-set-resource", es.pgtkSetResource_autogen, 2)
	ec.defSubr1(nil, "x-close-connection", es.xCloseConnection_autogen, 1)
	ec.defSubr1(nil, "x-close-connection", es.xCloseConnection_1_autogen, 1)
	ec.defSubr1(nil, "x-close-connection", es.xCloseConnection_2_autogen, 1)
	ec.defSubr1(nil, "x-create-frame", es.xCreateFrame_autogen, 1)
	ec.defSubr1(nil, "x-create-frame", es.xCreateFrame_1_autogen, 1)
	ec.defSubr1(nil, "x-create-frame", es.xCreateFrame_2_autogen, 1)
	ec.defSubr1(nil, "x-display-backing-store", es.xDisplayBackingStore_autogen, 0)
	ec.defSubr1(nil, "x-display-backing-store", es.xDisplayBackingStore_1_autogen, 0)
	ec.defSubr1(nil, "x-display-backing-store", es.xDisplayBackingStore_2_autogen, 0)
	ec.defSubr1(nil, "x-display-color-cells", es.xDisplayColorCells_autogen, 0)
	ec.defSubr1(nil, "x-display-color-cells", es.xDisplayColorCells_1_autogen, 0)
	ec.defSubr1(nil, "x-display-color-cells", es.xDisplayColorCells_2_autogen, 0)
	ec.defSubr1(nil, "x-display-grayscale-p", es.xDisplayGrayscaleP_autogen, 0)
	ec.defSubr1(nil, "x-display-grayscale-p", es.xDisplayGrayscaleP_1_autogen, 0)
	ec.defSubr1(nil, "x-display-grayscale-p", es.xDisplayGrayscaleP_2_autogen, 0)
	ec.defSubr0(nil, "x-display-list", es.xDisplayList_autogen)
	ec.defSubr0(nil, "x-display-list", es.xDisplayList_1_autogen)
	ec.defSubr0(nil, "x-display-list", es.xDisplayList_2_autogen)
	ec.defSubr1(nil, "x-display-mm-height", es.xDisplayMmHeight_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-height", es.xDisplayMmHeight_1_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-height", es.xDisplayMmHeight_2_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-width", es.xDisplayMmWidth_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-width", es.xDisplayMmWidth_1_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-width", es.xDisplayMmWidth_2_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-height", es.xDisplayPixelHeight_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-height", es.xDisplayPixelHeight_1_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-height", es.xDisplayPixelHeight_2_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-width", es.xDisplayPixelWidth_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-width", es.xDisplayPixelWidth_1_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-width", es.xDisplayPixelWidth_2_autogen, 0)
	ec.defSubr1(nil, "x-display-planes", es.xDisplayPlanes_autogen, 0)
	ec.defSubr1(nil, "x-display-planes", es.xDisplayPlanes_1_autogen, 0)
	ec.defSubr1(nil, "x-display-planes", es.xDisplayPlanes_2_autogen, 0)
	ec.defSubr1(nil, "x-display-save-under", es.xDisplaySaveUnder_autogen, 0)
	ec.defSubr1(nil, "x-display-save-under", es.xDisplaySaveUnder_1_autogen, 0)
	ec.defSubr1(nil, "x-display-save-under", es.xDisplaySaveUnder_2_autogen, 0)
	ec.defSubr1(nil, "x-display-screens", es.xDisplayScreens_autogen, 0)
	ec.defSubr1(nil, "x-display-screens", es.xDisplayScreens_1_autogen, 0)
	ec.defSubr1(nil, "x-display-screens", es.xDisplayScreens_2_autogen, 0)
	ec.defSubr1(nil, "x-display-visual-class", es.xDisplayVisualClass_autogen, 0)
	ec.defSubr1(nil, "x-display-visual-class", es.xDisplayVisualClass_1_autogen, 0)
	ec.defSubr1(nil, "x-display-visual-class", es.xDisplayVisualClass_2_autogen, 0)
	ec.defSubr2(nil, "x-export-frames", es.xExportFrames_autogen, 0)
	ec.defSubr2(nil, "x-export-frames", es.xExportFrames_1_autogen, 0)
	ec.defSubr5(nil, "x-file-dialog", es.xFileDialog_autogen, 2)
	ec.defSubr5(nil, "x-file-dialog", es.xFileDialog_1_autogen, 2)
	ec.defSubr5(nil, "x-file-dialog", es.xFileDialog_2_autogen, 2)
	ec.defSubr5(nil, "x-file-dialog", es.xFileDialog_3_autogen, 2)
	ec.defSubr1(nil, "x-gtk-debug", es.xGtkDebug_autogen, 1)
	ec.defSubr1(nil, "x-gtk-debug", es.xGtkDebug_1_autogen, 1)
	ec.defSubr0(nil, "x-hide-tip", es.xHideTip_autogen)
	ec.defSubr0(nil, "x-hide-tip", es.xHideTip_1_autogen)
	ec.defSubr0(nil, "x-hide-tip", es.xHideTip_2_autogen)
	ec.defSubr3(nil, "x-open-connection", es.xOpenConnection_autogen, 1)
	ec.defSubr3(nil, "x-open-connection", es.xOpenConnection_1_autogen, 1)
	ec.defSubr3(nil, "x-open-connection", es.xOpenConnection_2_autogen, 1)
	ec.defSubr2(nil, "x-select-font", es.xSelectFont_autogen, 0)
	ec.defSubr2(nil, "x-select-font", es.xSelectFont_1_autogen, 0)
	ec.defSubr2(nil, "x-select-font", es.xSelectFont_2_autogen, 0)
	ec.defSubr1(nil, "x-server-max-request-size", es.xServerMaxRequestSize_autogen, 0)
	ec.defSubr1(nil, "x-server-max-request-size", es.xServerMaxRequestSize_1_autogen, 0)
	ec.defSubr1(nil, "x-server-max-request-size", es.xServerMaxRequestSize_2_autogen, 0)
	ec.defSubr6(nil, "x-show-tip", es.xShowTip_autogen, 1)
	ec.defSubr6(nil, "x-show-tip", es.xShowTip_1_autogen, 1)
	ec.defSubr6(nil, "x-show-tip", es.xShowTip_2_autogen, 1)
	ec.defSubr2(nil, "xw-color-defined-p", es.xwColorDefinedP_autogen, 1)
	ec.defSubr2(nil, "xw-color-defined-p", es.xwColorDefinedP_1_autogen, 1)
	ec.defSubr2(nil, "xw-color-defined-p", es.xwColorDefinedP_2_autogen, 1)
	ec.defSubr2(nil, "xw-color-values", es.xwColorValues_autogen, 1)
	ec.defSubr2(nil, "xw-color-values", es.xwColorValues_1_autogen, 1)
	ec.defSubr2(nil, "xw-color-values", es.xwColorValues_2_autogen, 1)
	ec.defSubr1(nil, "xw-display-color-p", es.xwDisplayColorP_autogen, 0)
	ec.defSubr1(nil, "xw-display-color-p", es.xwDisplayColorP_1_autogen, 0)
	ec.defSubr1(nil, "xw-display-color-p", es.xwDisplayColorP_2_autogen, 0)
	ec.defSubr2(nil, "pgtk-use-im-context", es.pgtkUseImContext_autogen, 1)
	ec.defSubr0(nil, "menu-or-popup-active-p", es.menuOrPopupActiveP_autogen)
	ec.defSubr0(nil, "menu-or-popup-active-p", es.menuOrPopupActiveP_1_autogen)
	ec.defSubr0(nil, "menu-or-popup-active-p", es.menuOrPopupActiveP_2_autogen)
	ec.defSubr1(nil, "x-menu-bar-open-internal", es.xMenuBarOpenInternal_autogen, 0)
	ec.defSubr1(nil, "x-menu-bar-open-internal", es.xMenuBarOpenInternal_1_autogen, 0)
	ec.defSubr1(nil, "x-menu-bar-open-internal", es.xMenuBarOpenInternal_2_autogen, 0)
	ec.defSubr3(nil, "pgtk-disown-selection-internal", es.pgtkDisownSelectionInternal_autogen, 1)
	ec.defSubr3(nil, "pgtk-drop-finish", es.pgtkDropFinish_autogen, 3)
	ec.defSubr4(nil, "pgtk-get-selection-internal", es.pgtkGetSelectionInternal_autogen, 2)
	ec.defSubr3(nil, "pgtk-own-selection-internal", es.pgtkOwnSelectionInternal_autogen, 2)
	ec.defSubr2(nil, "pgtk-register-dnd-targets", es.pgtkRegisterDndTargets_autogen, 2)
	ec.defSubr2(nil, "pgtk-selection-exists-p", es.pgtkSelectionExistsP_autogen, 0)
	ec.defSubr2(nil, "pgtk-selection-owner-p", es.pgtkSelectionOwnerP_autogen, 0)
	ec.defSubr2(nil, "pgtk-update-drop-status", es.pgtkUpdateDropStatus_autogen, 2)
	ec.defSubr1(nil, "error-message-string", es.errorMessageString_autogen, 1)
	ec.defSubr1(nil, "external-debugging-output", es.externalDebuggingOutput_autogen, 1)
	ec.defSubr0(nil, "flush-standard-output", es.flushStandardOutput_autogen)
	ec.defSubr3(nil, "prin1", es.prin1_autogen, 1)
	ec.defSubr3(nil, "prin1-to-string", es.prin1ToString_autogen, 1)
	ec.defSubr2(nil, "princ", es.princ_autogen, 1)
	ec.defSubr2(nil, "print", es.print_autogen, 1)
	ec.defSubr1(nil, "print--preprocess", es.printPreprocess_autogen, 1)
	ec.defSubr2(nil, "redirect-debugging-output", es.redirectDebuggingOutput_autogen, 1)
	ec.defSubr2(nil, "terpri", es.terpri_autogen, 0)
	ec.defSubr2(nil, "write-char", es.writeChar_autogen, 1)
	ec.defSubr4(nil, "accept-process-output", es.acceptProcessOutput_autogen, 0)
	ec.defSubr2(nil, "continue-process", es.continueProcess_autogen, 0)
	ec.defSubr1(nil, "delete-process", es.deleteProcess_autogen, 0)
	ec.defSubr2(nil, "format-network-address", es.formatNetworkAddress_autogen, 1)
	ec.defSubr1(nil, "get-buffer-process", es.getBufferProcess_autogen, 1)
	ec.defSubr1(nil, "get-process", es.getProcess_autogen, 1)
	ec.defSubr2(nil, "internal-default-interrupt-process", es.internalDefaultInterruptProcess_autogen, 0)
	ec.defSubr2(nil, "internal-default-process-filter", es.internalDefaultProcessFilter_autogen, 2)
	ec.defSubr2(nil, "internal-default-process-sentinel", es.internalDefaultProcessSentinel_autogen, 2)
	ec.defSubr3(nil, "internal-default-signal-process", es.internalDefaultSignalProcess_autogen, 2)
	ec.defSubr2(nil, "interrupt-process", es.interruptProcess_autogen, 0)
	ec.defSubr2(nil, "kill-process", es.killProcess_autogen, 0)
	ec.defSubr0(nil, "list-system-processes", es.listSystemProcesses_autogen)
	ec.defSubrM(nil, "make-network-process", es.makeNetworkProcess_autogen, 0)
	ec.defSubrM(nil, "make-pipe-process", es.makePipeProcess_autogen, 0)
	ec.defSubrM(nil, "make-process", es.makeProcess_autogen, 0)
	ec.defSubrM(nil, "make-serial-process", es.makeSerialProcess_autogen, 0)
	ec.defSubr1(nil, "network-interface-info", es.networkInterfaceInfo_autogen, 1)
	ec.defSubr2(nil, "network-interface-list", es.networkInterfaceList_autogen, 0)
	ec.defSubr3(nil, "network-lookup-address-info", es.networkLookupAddressInfo_autogen, 1)
	ec.defSubr1(nil, "num-processors", es.numProcessors_autogen, 0)
	ec.defSubr1(nil, "process-attributes", es.processAttributes_autogen, 1)
	ec.defSubr1(nil, "process-buffer", es.processBuffer_autogen, 1)
	ec.defSubr1(nil, "process-coding-system", es.processCodingSystem_autogen, 1)
	ec.defSubr1(nil, "process-command", es.processCommand_autogen, 1)
	ec.defSubr1(nil, "process-connection", es.processConnection_autogen, 1)
	ec.defSubr3(nil, "process-contact", es.processContact_autogen, 1)
	ec.defSubr1(nil, "process-datagram-address", es.processDatagramAddress_autogen, 1)
	ec.defSubr1(nil, "process-exit-status", es.processExitStatus_autogen, 1)
	ec.defSubr1(nil, "process-filter", es.processFilter_autogen, 1)
	ec.defSubr1(nil, "process-id", es.processId_autogen, 1)
	ec.defSubr1(nil, "process-inherit-coding-system-flag", es.processInheritCodingSystemFlag_autogen, 1)
	ec.defSubr0(nil, "process-list", es.processList_autogen)
	ec.defSubr1(nil, "process-mark", es.processMark_autogen, 1)
	ec.defSubr1(nil, "process-name", es.processName_autogen, 1)
	ec.defSubr1(nil, "process-plist", es.processPlist_autogen, 1)
	ec.defSubr1(nil, "process-query-on-exit-flag", es.processQueryOnExitFlag_autogen, 1)
	ec.defSubr1(nil, "process-running-child-p", es.processRunningChildP_autogen, 0)
	ec.defSubr1(nil, "process-send-eof", es.processSendEof_autogen, 0)
	ec.defSubr3(nil, "process-send-region", es.processSendRegion_autogen, 3)
	ec.defSubr2(nil, "process-send-string", es.processSendString_autogen, 2)
	ec.defSubr1(nil, "process-sentinel", es.processSentinel_autogen, 1)
	ec.defSubr1(nil, "process-status", es.processStatus_autogen, 1)
	ec.defSubr1(nil, "process-thread", es.processThread_autogen, 1)
	ec.defSubr2(nil, "process-tty-name", es.processTtyName_autogen, 1)
	ec.defSubr1(nil, "process-type", es.processType_autogen, 1)
	ec.defSubr1(nil, "processp", es.processp_autogen, 1)
	ec.defSubr2(nil, "quit-process", es.quitProcess_autogen, 0)
	ec.defSubrM(nil, "serial-process-configure", es.serialProcessConfigure_autogen, 0)
	ec.defSubr4(nil, "set-network-process-option", es.setNetworkProcessOption_autogen, 3)
	ec.defSubr2(nil, "set-process-buffer", es.setProcessBuffer_autogen, 2)
	ec.defSubr3(nil, "set-process-coding-system", es.setProcessCodingSystem_autogen, 1)
	ec.defSubr2(nil, "set-process-datagram-address", es.setProcessDatagramAddress_autogen, 2)
	ec.defSubr2(nil, "set-process-filter", es.setProcessFilter_autogen, 2)
	ec.defSubr2(nil, "set-process-inherit-coding-system-flag", es.setProcessInheritCodingSystemFlag_autogen, 2)
	ec.defSubr2(nil, "set-process-plist", es.setProcessPlist_autogen, 2)
	ec.defSubr2(nil, "set-process-query-on-exit-flag", es.setProcessQueryOnExitFlag_autogen, 2)
	ec.defSubr2(nil, "set-process-sentinel", es.setProcessSentinel_autogen, 2)
	ec.defSubr2(nil, "set-process-thread", es.setProcessThread_autogen, 2)
	ec.defSubr3(nil, "set-process-window-size", es.setProcessWindowSize_autogen, 3)
	ec.defSubr0(nil, "signal-names", es.signalNames_autogen)
	ec.defSubr3(nil, "signal-process", es.signalProcess_autogen, 2)
	ec.defSubr2(nil, "stop-process", es.stopProcess_autogen, 0)
	ec.defSubr0(nil, "waiting-for-user-input-p", es.waitingForUserInputP_autogen)
	ec.defSubr2(nil, "function-equal", es.functionEqual_autogen, 2)
	ec.defSubr0(nil, "profiler-cpu-log", es.profilerCpuLog_autogen)
	ec.defSubr0(nil, "profiler-cpu-running-p", es.profilerCpuRunningP_autogen)
	ec.defSubr1(nil, "profiler-cpu-start", es.profilerCpuStart_autogen, 1)
	ec.defSubr0(nil, "profiler-cpu-stop", es.profilerCpuStop_autogen)
	ec.defSubr0(nil, "profiler-memory-log", es.profilerMemoryLog_autogen)
	ec.defSubr0(nil, "profiler-memory-running-p", es.profilerMemoryRunningP_autogen)
	ec.defSubr0(nil, "profiler-memory-start", es.profilerMemoryStart_autogen)
	ec.defSubr0(nil, "profiler-memory-stop", es.profilerMemoryStop_autogen)
	ec.defSubr2(nil, "looking-at", es.lookingAt_autogen, 1)
	ec.defSubr1(nil, "match-beginning", es.matchBeginning_autogen, 1)
	ec.defSubr3(nil, "match-data", es.matchData_autogen, 0)
	ec.defSubr1(nil, "match-data--translate", es.matchDataTranslate_autogen, 1)
	ec.defSubr1(nil, "match-end", es.matchEnd_autogen, 1)
	ec.defSubr1(nil, "newline-cache-check", es.newlineCacheCheck_autogen, 0)
	ec.defSubr2(nil, "posix-looking-at", es.posixLookingAt_autogen, 1)
	ec.defSubr4(nil, "posix-search-backward", es.posixSearchBackward_autogen, 1)
	ec.defSubr4(nil, "posix-search-forward", es.posixSearchForward_autogen, 1)
	ec.defSubr4(nil, "posix-string-match", es.posixStringMatch_autogen, 2)
	ec.defSubr4(nil, "re-search-backward", es.reSearchBackward_autogen, 1)
	ec.defSubr4(nil, "re-search-forward", es.reSearchForward_autogen, 1)
	ec.defSubr1(nil, "regexp-quote", es.regexpQuote_autogen, 1)
	ec.defSubr5(nil, "replace-match", es.replaceMatch_autogen, 1)
	ec.defSubr4(nil, "search-backward", es.searchBackward_autogen, 1)
	ec.defSubr4(nil, "search-forward", es.searchForward_autogen, 1)
	ec.defSubr2(nil, "set-match-data", es.setMatchData_autogen, 1)
	ec.defSubr4(nil, "string-match", es.stringMatch_autogen, 2)
	ec.defSubr1(nil, "play-sound-internal", es.playSoundInternal_autogen, 1)
	ec.defSubr0(nil, "sqlite-available-p", es.sqliteAvailableP_autogen)
	ec.defSubr1(nil, "sqlite-close", es.sqliteClose_autogen, 1)
	ec.defSubr1(nil, "sqlite-columns", es.sqliteColumns_autogen, 1)
	ec.defSubr1(nil, "sqlite-commit", es.sqliteCommit_autogen, 1)
	ec.defSubr3(nil, "sqlite-execute", es.sqliteExecute_autogen, 2)
	ec.defSubr1(nil, "sqlite-finalize", es.sqliteFinalize_autogen, 1)
	ec.defSubr2(nil, "sqlite-load-extension", es.sqliteLoadExtension_autogen, 2)
	ec.defSubr1(nil, "sqlite-more-p", es.sqliteMoreP_autogen, 1)
	ec.defSubr1(nil, "sqlite-next", es.sqliteNext_autogen, 1)
	ec.defSubr1(nil, "sqlite-open", es.sqliteOpen_autogen, 0)
	ec.defSubr2(nil, "sqlite-pragma", es.sqlitePragma_autogen, 2)
	ec.defSubr1(nil, "sqlite-rollback", es.sqliteRollback_autogen, 1)
	ec.defSubr4(nil, "sqlite-select", es.sqliteSelect_autogen, 2)
	ec.defSubr1(nil, "sqlite-transaction", es.sqliteTransaction_autogen, 1)
	ec.defSubr0(nil, "sqlite-version", es.sqliteVersion_autogen)
	ec.defSubr1(nil, "sqlitep", es.sqlitep_autogen, 1)
	ec.defSubr0(nil, "backward-prefix-chars", es.backwardPrefixChars_autogen)
	ec.defSubr1(nil, "char-syntax", es.charSyntax_autogen, 1)
	ec.defSubr1(nil, "copy-syntax-table", es.copySyntaxTable_autogen, 0)
	ec.defSubr1(nil, "forward-comment", es.forwardComment_autogen, 1)
	ec.defSubr1(nil, "forward-word", es.forwardWord_autogen, 0)
	ec.defSubr1(nil, "internal-describe-syntax-value", es.internalDescribeSyntaxValue_autogen, 1)
	ec.defSubr1(nil, "matching-paren", es.matchingParen_autogen, 1)
	ec.defSubr3(nil, "modify-syntax-entry", es.modifySyntaxEntry_autogen, 2)
	ec.defSubr6(nil, "parse-partial-sexp", es.parsePartialSexp_autogen, 2)
	ec.defSubr3(nil, "scan-lists", es.scanLists_autogen, 3)
	ec.defSubr2(nil, "scan-sexps", es.scanSexps_autogen, 2)
	ec.defSubr1(nil, "set-syntax-table", es.setSyntaxTable_autogen, 1)
	ec.defSubr2(nil, "skip-chars-backward", es.skipCharsBackward_autogen, 1)
	ec.defSubr2(nil, "skip-chars-forward", es.skipCharsForward_autogen, 1)
	ec.defSubr2(nil, "skip-syntax-backward", es.skipSyntaxBackward_autogen, 1)
	ec.defSubr2(nil, "skip-syntax-forward", es.skipSyntaxForward_autogen, 1)
	ec.defSubr0(nil, "standard-syntax-table", es.standardSyntaxTable_autogen)
	ec.defSubr1(nil, "string-to-syntax", es.stringToSyntax_autogen, 1)
	ec.defSubr1(nil, "syntax-class-to-char", es.syntaxClassToChar_autogen, 1)
	ec.defSubr0(nil, "syntax-table", es.syntaxTable_autogen)
	ec.defSubr1(nil, "syntax-table-p", es.syntaxTableP_autogen, 1)
	ec.defSubr0(nil, "get-internal-run-time", es.getInternalRunTime_autogen)
	ec.defSubr1(nil, "controlling-tty-p", es.controllingTtyP_autogen, 0)
	ec.defSubr0(nil, "gpm-mouse-start", es.gpmMouseStart_autogen)
	ec.defSubr0(nil, "gpm-mouse-stop", es.gpmMouseStop_autogen)
	ec.defSubr1(nil, "resume-tty", es.resumeTty_autogen, 0)
	ec.defSubr1(nil, "suspend-tty", es.suspendTty_autogen, 0)
	ec.defSubr1(nil, "tty--output-buffer-size", es.ttyOutputBufferSize_autogen, 0)
	ec.defSubr2(nil, "tty--set-output-buffer-size", es.ttySetOutputBufferSize_autogen, 1)
	ec.defSubr1(nil, "tty-display-color-cells", es.ttyDisplayColorCells_autogen, 0)
	ec.defSubr1(nil, "tty-display-color-p", es.ttyDisplayColorP_autogen, 0)
	ec.defSubr1(nil, "tty-no-underline", es.ttyNoUnderline_autogen, 0)
	ec.defSubr1(nil, "tty-top-frame", es.ttyTopFrame_autogen, 0)
	ec.defSubr1(nil, "tty-type", es.ttyType_autogen, 0)
	ec.defSubr2(nil, "delete-terminal", es.deleteTerminal_autogen, 0)
	ec.defSubr1(nil, "frame-terminal", es.frameTerminal_autogen, 0)
	ec.defSubr3(nil, "set-terminal-parameter", es.setTerminalParameter_autogen, 3)
	ec.defSubr0(nil, "terminal-list", es.terminalList_autogen)
	ec.defSubr1(nil, "terminal-live-p", es.terminalLiveP_autogen, 1)
	ec.defSubr1(nil, "terminal-name", es.terminalName_autogen, 0)
	ec.defSubr2(nil, "terminal-parameter", es.terminalParameter_autogen, 2)
	ec.defSubr1(nil, "terminal-parameters", es.terminalParameters_autogen, 0)
	ec.defSubr2(nil, "set-text-conversion-style", es.setTextConversionStyle_autogen, 1)
	ec.defSubr5(nil, "add-face-text-property", es.addFaceTextProperty_autogen, 3)
	ec.defSubr4(nil, "add-text-properties", es.addTextProperties_autogen, 3)
	ec.defSubr3(nil, "get-char-property", es.getCharProperty_autogen, 2)
	ec.defSubr3(nil, "get-char-property-and-overlay", es.getCharPropertyAndOverlay_autogen, 2)
	ec.defSubr3(nil, "get-text-property", es.getTextProperty_autogen, 2)
	ec.defSubr2(nil, "next-char-property-change", es.nextCharPropertyChange_autogen, 1)
	ec.defSubr3(nil, "next-property-change", es.nextPropertyChange_autogen, 1)
	ec.defSubr4(nil, "next-single-char-property-change", es.nextSingleCharPropertyChange_autogen, 2)
	ec.defSubr4(nil, "next-single-property-change", es.nextSinglePropertyChange_autogen, 2)
	ec.defSubr2(nil, "previous-char-property-change", es.previousCharPropertyChange_autogen, 1)
	ec.defSubr3(nil, "previous-property-change", es.previousPropertyChange_autogen, 1)
	ec.defSubr4(nil, "previous-single-char-property-change", es.previousSingleCharPropertyChange_autogen, 2)
	ec.defSubr4(nil, "previous-single-property-change", es.previousSinglePropertyChange_autogen, 2)
	ec.defSubr5(nil, "put-text-property", es.putTextProperty_autogen, 4)
	ec.defSubr4(nil, "remove-list-of-text-properties", es.removeListOfTextProperties_autogen, 3)
	ec.defSubr4(nil, "remove-text-properties", es.removeTextProperties_autogen, 3)
	ec.defSubr4(nil, "set-text-properties", es.setTextProperties_autogen, 3)
	ec.defSubr2(nil, "text-properties-at", es.textPropertiesAt_autogen, 1)
	ec.defSubr5(nil, "text-property-any", es.textPropertyAny_autogen, 4)
	ec.defSubr5(nil, "text-property-not-all", es.textPropertyNotAll_autogen, 4)
	ec.defSubr0(nil, "all-threads", es.allThreads_autogen)
	ec.defSubr1(nil, "condition-mutex", es.conditionMutex_autogen, 1)
	ec.defSubr1(nil, "condition-name", es.conditionName_autogen, 1)
	ec.defSubr2(nil, "condition-notify", es.conditionNotify_autogen, 1)
	ec.defSubr1(nil, "condition-wait", es.conditionWait_autogen, 1)
	ec.defSubr0(nil, "current-thread", es.currentThread_autogen)
	ec.defSubr2(nil, "make-condition-variable", es.makeConditionVariable_autogen, 1)
	ec.defSubr1(nil, "make-mutex", es.makeMutex_autogen, 0)
	ec.defSubr2(nil, "make-thread", es.makeThread_autogen, 1)
	ec.defSubr1(nil, "mutex-lock", es.mutexLock_autogen, 1)
	ec.defSubr1(nil, "mutex-name", es.mutexName_autogen, 1)
	ec.defSubr1(nil, "mutex-unlock", es.mutexUnlock_autogen, 1)
	ec.defSubr1(nil, "thread--blocker", es.threadBlocker_autogen, 1)
	ec.defSubr1(nil, "thread-join", es.threadJoin_autogen, 1)
	ec.defSubr1(nil, "thread-last-error", es.threadLastError_autogen, 0)
	ec.defSubr1(nil, "thread-live-p", es.threadLiveP_autogen, 1)
	ec.defSubr1(nil, "thread-name", es.threadName_autogen, 1)
	ec.defSubr3(nil, "thread-signal", es.threadSignal_autogen, 3)
	ec.defSubr0(nil, "thread-yield", es.threadYield_autogen)
	ec.defSubr0(nil, "current-cpu-time", es.currentCpuTime_autogen)
	ec.defSubr0(nil, "current-time", es.currentTime_autogen)
	ec.defSubr2(nil, "current-time-string", es.currentTimeString_autogen, 0)
	ec.defSubr2(nil, "current-time-zone", es.currentTimeZone_autogen, 0)
	ec.defSubr3(nil, "decode-time", es.decodeTime_autogen, 0)
	ec.defSubrM(nil, "encode-time", es.encodeTime_autogen, 1)
	ec.defSubr1(nil, "float-time", es.floatTime_autogen, 0)
	ec.defSubr3(nil, "format-time-string", es.formatTimeString_autogen, 1)
	ec.defSubr1(nil, "set-time-zone-rule", es.setTimeZoneRule_autogen, 1)
	ec.defSubr2(nil, "time-add", es.timeAdd_autogen, 2)
	ec.defSubr2(nil, "time-convert", es.timeConvert_autogen, 1)
	ec.defSubr2(nil, "time-equal-p", es.timeEqualP_autogen, 2)
	ec.defSubr2(nil, "time-less-p", es.timeLessP_autogen, 2)
	ec.defSubr2(nil, "time-subtract", es.timeSubtract_autogen, 2)
	ec.defSubr0(nil, "treesit-available-p", es.treesitAvailableP_autogen)
	ec.defSubr1(nil, "treesit-compiled-query-p", es.treesitCompiledQueryP_autogen, 1)
	ec.defSubr4(nil, "treesit-induce-sparse-tree", es.treesitInduceSparseTree_autogen, 2)
	ec.defSubr1(nil, "treesit-language-abi-version", es.treesitLanguageAbiVersion_autogen, 0)
	ec.defSubr2(nil, "treesit-language-available-p", es.treesitLanguageAvailableP_autogen, 1)
	ec.defSubr1(nil, "treesit-library-abi-version", es.treesitLibraryAbiVersion_autogen, 0)
	ec.defSubr2(nil, "treesit-node-check", es.treesitNodeCheck_autogen, 2)
	ec.defSubr3(nil, "treesit-node-child", es.treesitNodeChild_autogen, 2)
	ec.defSubr2(nil, "treesit-node-child-by-field-name", es.treesitNodeChildByFieldName_autogen, 2)
	ec.defSubr2(nil, "treesit-node-child-count", es.treesitNodeChildCount_autogen, 1)
	ec.defSubr4(nil, "treesit-node-descendant-for-range", es.treesitNodeDescendantForRange_autogen, 3)
	ec.defSubr1(nil, "treesit-node-end", es.treesitNodeEnd_autogen, 1)
	ec.defSubr2(nil, "treesit-node-eq", es.treesitNodeEq_autogen, 2)
	ec.defSubr2(nil, "treesit-node-field-name-for-child", es.treesitNodeFieldNameForChild_autogen, 2)
	ec.defSubr3(nil, "treesit-node-first-child-for-pos", es.treesitNodeFirstChildForPos_autogen, 2)
	ec.defSubr2(nil, "treesit-node-match-p", es.treesitNodeMatchP_autogen, 2)
	ec.defSubr2(nil, "treesit-node-next-sibling", es.treesitNodeNextSibling_autogen, 1)
	ec.defSubr1(nil, "treesit-node-p", es.treesitNodeP_autogen, 1)
	ec.defSubr1(nil, "treesit-node-parent", es.treesitNodeParent_autogen, 1)
	ec.defSubr1(nil, "treesit-node-parser", es.treesitNodeParser_autogen, 1)
	ec.defSubr2(nil, "treesit-node-prev-sibling", es.treesitNodePrevSibling_autogen, 1)
	ec.defSubr1(nil, "treesit-node-start", es.treesitNodeStart_autogen, 1)
	ec.defSubr1(nil, "treesit-node-string", es.treesitNodeString_autogen, 1)
	ec.defSubr1(nil, "treesit-node-type", es.treesitNodeType_autogen, 1)
	ec.defSubr2(nil, "treesit-parser-add-notifier", es.treesitParserAddNotifier_autogen, 2)
	ec.defSubr1(nil, "treesit-parser-buffer", es.treesitParserBuffer_autogen, 1)
	ec.defSubr3(nil, "treesit-parser-create", es.treesitParserCreate_autogen, 1)
	ec.defSubr1(nil, "treesit-parser-delete", es.treesitParserDelete_autogen, 1)
	ec.defSubr1(nil, "treesit-parser-included-ranges", es.treesitParserIncludedRanges_autogen, 1)
	ec.defSubr1(nil, "treesit-parser-language", es.treesitParserLanguage_autogen, 1)
	ec.defSubr1(nil, "treesit-parser-list", es.treesitParserList_autogen, 0)
	ec.defSubr1(nil, "treesit-parser-notifiers", es.treesitParserNotifiers_autogen, 1)
	ec.defSubr1(nil, "treesit-parser-p", es.treesitParserP_autogen, 1)
	ec.defSubr2(nil, "treesit-parser-remove-notifier", es.treesitParserRemoveNotifier_autogen, 2)
	ec.defSubr1(nil, "treesit-parser-root-node", es.treesitParserRootNode_autogen, 1)
	ec.defSubr2(nil, "treesit-parser-set-included-ranges", es.treesitParserSetIncludedRanges_autogen, 2)
	ec.defSubr1(nil, "treesit-pattern-expand", es.treesitPatternExpand_autogen, 1)
	ec.defSubr5(nil, "treesit-query-capture", es.treesitQueryCapture_autogen, 2)
	ec.defSubr3(nil, "treesit-query-compile", es.treesitQueryCompile_autogen, 2)
	ec.defSubr1(nil, "treesit-query-expand", es.treesitQueryExpand_autogen, 1)
	ec.defSubr1(nil, "treesit-query-language", es.treesitQueryLanguage_autogen, 1)
	ec.defSubr1(nil, "treesit-query-p", es.treesitQueryP_autogen, 1)
	ec.defSubr4(nil, "treesit-search-forward", es.treesitSearchForward_autogen, 2)
	ec.defSubr5(nil, "treesit-search-subtree", es.treesitSearchSubtree_autogen, 2)
	ec.defSubr1(nil, "treesit-subtree-stat", es.treesitSubtreeStat_autogen, 1)
	ec.defSubr0(nil, "undo-boundary", es.undoBoundary_autogen)
	ec.defSubr1(nil, "w16-get-clipboard-data", es.w16GetClipboardData_autogen, 0)
	ec.defSubr2(nil, "w16-selection-exists-p", es.w16SelectionExistsP_autogen, 0)
	ec.defSubr2(nil, "w16-set-clipboard-data", es.w16SetClipboardData_autogen, 1)
	ec.defSubr0(nil, "get-screen-color", es.getScreenColor_autogen)
	ec.defSubr1(nil, "set-cursor-size", es.setCursorSize_autogen, 1)
	ec.defSubr2(nil, "set-screen-color", es.setScreenColor_autogen, 2)
	ec.defSubr0(nil, "w32-battery-status", es.w32BatteryStatus_autogen)
	ec.defSubr0(nil, "default-printer-name", es.defaultPrinterName_autogen)
	ec.defSubr1(nil, "set-message-beep", es.setMessageBeep_autogen, 1)
	ec.defSubr1(nil, "system-move-file-to-trash", es.systemMoveFileToTrash_autogen, 1)
	ec.defSubr0(nil, "w32--menu-bar-in-use", es.w32MenuBarInUse_autogen)
	ec.defSubr4(nil, "w32-define-rgb-color", es.w32DefineRgbColor_autogen, 4)
	ec.defSubr1(nil, "w32-display-monitor-attributes-list", es.w32DisplayMonitorAttributesList_autogen, 0)
	ec.defSubr2(nil, "w32-frame-edges", es.w32FrameEdges_autogen, 0)
	ec.defSubr1(nil, "w32-frame-geometry", es.w32FrameGeometry_autogen, 0)
	ec.defSubr1(nil, "w32-frame-list-z-order", es.w32FrameListZOrder_autogen, 0)
	ec.defSubr3(nil, "w32-frame-restack", es.w32FrameRestack_autogen, 2)
	ec.defSubr0(nil, "w32-get-ime-open-status", es.w32GetImeOpenStatus_autogen)
	ec.defSubr0(nil, "w32-mouse-absolute-pixel-position", es.w32MouseAbsolutePixelPosition_autogen)
	ec.defSubr1(nil, "w32-notification-close", es.w32NotificationClose_autogen, 1)
	ec.defSubrM(nil, "w32-notification-notify", es.w32NotificationNotify_autogen, 0)
	ec.defSubr3(nil, "w32-read-registry", es.w32ReadRegistry_autogen, 3)
	ec.defSubr1(nil, "w32-reconstruct-hot-key", es.w32ReconstructHotKey_autogen, 1)
	ec.defSubr1(nil, "w32-register-hot-key", es.w32RegisterHotKey_autogen, 1)
	ec.defSubr0(nil, "w32-registered-hot-keys", es.w32RegisteredHotKeys_autogen)
	ec.defSubr2(nil, "w32-send-sys-command", es.w32SendSysCommand_autogen, 1)
	ec.defSubr1(nil, "w32-set-ime-open-status", es.w32SetImeOpenStatus_autogen, 1)
	ec.defSubr2(nil, "w32-set-mouse-absolute-pixel-position", es.w32SetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr1(nil, "w32-set-wallpaper", es.w32SetWallpaper_autogen, 1)
	ec.defSubr4(nil, "w32-shell-execute", es.w32ShellExecute_autogen, 2)
	ec.defSubr2(nil, "w32-toggle-lock-key", es.w32ToggleLockKey_autogen, 1)
	ec.defSubr1(nil, "w32-unregister-hot-key", es.w32UnregisterHotKey_autogen, 1)
	ec.defSubr2(nil, "w32-window-exists-p", es.w32WindowExistsP_autogen, 2)
	ec.defSubr6(nil, "x-change-window-property", es.xChangeWindowProperty_autogen, 2)
	ec.defSubr7(nil, "x-change-window-property", es.xChangeWindowProperty_1_autogen, 2)
	ec.defSubr2(nil, "x-delete-window-property", es.xDeleteWindowProperty_autogen, 1)
	ec.defSubr3(nil, "x-delete-window-property", es.xDeleteWindowProperty_1_autogen, 1)
	ec.defSubr1(nil, "x-server-vendor", es.xServerVendor_autogen, 0)
	ec.defSubr1(nil, "x-server-vendor", es.xServerVendor_1_autogen, 0)
	ec.defSubr1(nil, "x-server-version", es.xServerVersion_autogen, 0)
	ec.defSubr1(nil, "x-server-version", es.xServerVersion_1_autogen, 0)
	ec.defSubr2(nil, "x-synchronize", es.xSynchronize_autogen, 1)
	ec.defSubr2(nil, "x-synchronize", es.xSynchronize_1_autogen, 1)
	ec.defSubr6(nil, "x-window-property", es.xWindowProperty_autogen, 1)
	ec.defSubr6(nil, "x-window-property", es.xWindowProperty_1_autogen, 1)
	ec.defSubr3(nil, "w32notify-add-watch", es.w32notifyAddWatch_autogen, 3)
	ec.defSubr1(nil, "w32notify-rm-watch", es.w32notifyRmWatch_autogen, 1)
	ec.defSubr1(nil, "w32notify-valid-p", es.w32notifyValidP_autogen, 1)
	ec.defSubr1(nil, "w32-application-type", es.w32ApplicationType_autogen, 1)
	ec.defSubr1(nil, "w32-get-codepage-charset", es.w32GetCodepageCharset_autogen, 1)
	ec.defSubr0(nil, "w32-get-console-codepage", es.w32GetConsoleCodepage_autogen)
	ec.defSubr0(nil, "w32-get-console-output-codepage", es.w32GetConsoleOutputCodepage_autogen)
	ec.defSubr0(nil, "w32-get-current-locale-id", es.w32GetCurrentLocaleId_autogen)
	ec.defSubr1(nil, "w32-get-default-locale-id", es.w32GetDefaultLocaleId_autogen, 0)
	ec.defSubr0(nil, "w32-get-keyboard-layout", es.w32GetKeyboardLayout_autogen)
	ec.defSubr2(nil, "w32-get-locale-info", es.w32GetLocaleInfo_autogen, 1)
	ec.defSubr0(nil, "w32-get-valid-codepages", es.w32GetValidCodepages_autogen)
	ec.defSubr0(nil, "w32-get-valid-keyboard-layouts", es.w32GetValidKeyboardLayouts_autogen)
	ec.defSubr0(nil, "w32-get-valid-locale-ids", es.w32GetValidLocaleIds_autogen)
	ec.defSubr1(nil, "w32-has-winsock", es.w32HasWinsock_autogen, 0)
	ec.defSubr1(nil, "w32-long-file-name", es.w32LongFileName_autogen, 1)
	ec.defSubr1(nil, "w32-set-console-codepage", es.w32SetConsoleCodepage_autogen, 1)
	ec.defSubr1(nil, "w32-set-console-output-codepage", es.w32SetConsoleOutputCodepage_autogen, 1)
	ec.defSubr1(nil, "w32-set-current-locale", es.w32SetCurrentLocale_autogen, 1)
	ec.defSubr1(nil, "w32-set-keyboard-layout", es.w32SetKeyboardLayout_autogen, 1)
	ec.defSubr2(nil, "w32-set-process-priority", es.w32SetProcessPriority_autogen, 2)
	ec.defSubr1(nil, "w32-short-file-name", es.w32ShortFileName_autogen, 1)
	ec.defSubr0(nil, "w32-unload-winsock", es.w32UnloadWinsock_autogen)
	ec.defSubr1(nil, "w32-get-clipboard-data", es.w32GetClipboardData_autogen, 0)
	ec.defSubr2(nil, "w32-selection-exists-p", es.w32SelectionExistsP_autogen, 0)
	ec.defSubr2(nil, "w32-selection-targets", es.w32SelectionTargets_autogen, 0)
	ec.defSubr2(nil, "w32-set-clipboard-data", es.w32SetClipboardData_autogen, 1)
	ec.defSubr2(nil, "coordinates-in-window-p", es.coordinatesInWindowP_autogen, 2)
	ec.defSubr1(nil, "current-window-configuration", es.currentWindowConfiguration_autogen, 0)
	ec.defSubr2(nil, "delete-other-windows-internal", es.deleteOtherWindowsInternal_autogen, 0)
	ec.defSubr1(nil, "delete-window-internal", es.deleteWindowInternal_autogen, 1)
	ec.defSubr1(nil, "force-window-update", es.forceWindowUpdate_autogen, 0)
	ec.defSubr1(nil, "frame-first-window", es.frameFirstWindow_autogen, 0)
	ec.defSubr1(nil, "frame-old-selected-window", es.frameOldSelectedWindow_autogen, 0)
	ec.defSubr1(nil, "frame-root-window", es.frameRootWindow_autogen, 0)
	ec.defSubr1(nil, "frame-selected-window", es.frameSelectedWindow_autogen, 0)
	ec.defSubr2(nil, "get-buffer-window", es.getBufferWindow_autogen, 0)
	ec.defSubr0(nil, "minibuffer-selected-window", es.minibufferSelectedWindow_autogen)
	ec.defSubr1(nil, "minibuffer-window", es.minibufferWindow_autogen, 0)
	ec.defSubr1(nil, "move-to-window-line", es.moveToWindowLine_autogen, 1)
	ec.defSubr3(nil, "next-window", es.nextWindow_autogen, 0)
	ec.defSubr0(nil, "old-selected-window", es.oldSelectedWindow_autogen)
	ec.defSubr0(nil, "other-window-for-scrolling", es.otherWindowForScrolling_autogen)
	ec.defSubr3(nil, "pos-visible-in-window-p", es.posVisibleInWindowP_autogen, 0)
	ec.defSubr3(nil, "previous-window", es.previousWindow_autogen, 0)
	ec.defSubr2(nil, "recenter", es.recenter_autogen, 0)
	ec.defSubr1(nil, "resize-mini-window-internal", es.resizeMiniWindowInternal_autogen, 1)
	ec.defSubr1(nil, "run-window-configuration-change-hook", es.runWindowConfigurationChangeHook_autogen, 0)
	ec.defSubr1(nil, "run-window-scroll-functions", es.runWindowScrollFunctions_autogen, 0)
	ec.defSubr1(nil, "scroll-down", es.scrollDown_autogen, 0)
	ec.defSubr2(nil, "scroll-left", es.scrollLeft_autogen, 0)
	ec.defSubr2(nil, "scroll-right", es.scrollRight_autogen, 0)
	ec.defSubr1(nil, "scroll-up", es.scrollUp_autogen, 0)
	ec.defSubr2(nil, "select-window", es.selectWindow_autogen, 1)
	ec.defSubr0(nil, "selected-window", es.selectedWindow_autogen)
	ec.defSubr3(nil, "set-frame-selected-window", es.setFrameSelectedWindow_autogen, 2)
	ec.defSubr3(nil, "set-window-buffer", es.setWindowBuffer_autogen, 2)
	ec.defSubr2(nil, "set-window-combination-limit", es.setWindowCombinationLimit_autogen, 2)
	ec.defSubr3(nil, "set-window-configuration", es.setWindowConfiguration_autogen, 1)
	ec.defSubr2(nil, "set-window-dedicated-p", es.setWindowDedicatedP_autogen, 2)
	ec.defSubr2(nil, "set-window-display-table", es.setWindowDisplayTable_autogen, 2)
	ec.defSubr5(nil, "set-window-fringes", es.setWindowFringes_autogen, 2)
	ec.defSubr2(nil, "set-window-hscroll", es.setWindowHscroll_autogen, 2)
	ec.defSubr3(nil, "set-window-margins", es.setWindowMargins_autogen, 2)
	ec.defSubr2(nil, "set-window-new-normal", es.setWindowNewNormal_autogen, 1)
	ec.defSubr3(nil, "set-window-new-pixel", es.setWindowNewPixel_autogen, 2)
	ec.defSubr3(nil, "set-window-new-total", es.setWindowNewTotal_autogen, 2)
	ec.defSubr2(nil, "set-window-next-buffers", es.setWindowNextBuffers_autogen, 2)
	ec.defSubr3(nil, "set-window-parameter", es.setWindowParameter_autogen, 3)
	ec.defSubr2(nil, "set-window-point", es.setWindowPoint_autogen, 2)
	ec.defSubr2(nil, "set-window-prev-buffers", es.setWindowPrevBuffers_autogen, 2)
	ec.defSubr6(nil, "set-window-scroll-bars", es.setWindowScrollBars_autogen, 1)
	ec.defSubr3(nil, "set-window-start", es.setWindowStart_autogen, 2)
	ec.defSubr4(nil, "set-window-vscroll", es.setWindowVscroll_autogen, 2)
	ec.defSubr4(nil, "split-window-internal", es.splitWindowInternal_autogen, 4)
	ec.defSubr3(nil, "window-at", es.windowAt_autogen, 2)
	ec.defSubr2(nil, "window-body-height", es.windowBodyHeight_autogen, 0)
	ec.defSubr2(nil, "window-body-width", es.windowBodyWidth_autogen, 0)
	ec.defSubr1(nil, "window-bottom-divider-width", es.windowBottomDividerWidth_autogen, 0)
	ec.defSubr1(nil, "window-buffer", es.windowBuffer_autogen, 0)
	ec.defSubr1(nil, "window-bump-use-time", es.windowBumpUseTime_autogen, 0)
	ec.defSubr1(nil, "window-combination-limit", es.windowCombinationLimit_autogen, 1)
	ec.defSubr2(nil, "window-configuration-equal-p", es.windowConfigurationEqualP_autogen, 2)
	ec.defSubr1(nil, "window-configuration-frame", es.windowConfigurationFrame_autogen, 1)
	ec.defSubr1(nil, "window-configuration-p", es.windowConfigurationP_autogen, 1)
	ec.defSubr1(nil, "window-dedicated-p", es.windowDedicatedP_autogen, 0)
	ec.defSubr1(nil, "window-display-table", es.windowDisplayTable_autogen, 0)
	ec.defSubr2(nil, "window-end", es.windowEnd_autogen, 0)
	ec.defSubr1(nil, "window-frame", es.windowFrame_autogen, 0)
	ec.defSubr1(nil, "window-fringes", es.windowFringes_autogen, 0)
	ec.defSubr1(nil, "window-header-line-height", es.windowHeaderLineHeight_autogen, 0)
	ec.defSubr1(nil, "window-hscroll", es.windowHscroll_autogen, 0)
	ec.defSubr1(nil, "window-left-child", es.windowLeftChild_autogen, 0)
	ec.defSubr1(nil, "window-left-column", es.windowLeftColumn_autogen, 0)
	ec.defSubr2(nil, "window-line-height", es.windowLineHeight_autogen, 0)
	ec.defSubr6(nil, "window-lines-pixel-dimensions", es.windowLinesPixelDimensions_autogen, 0)
	ec.defSubr3(nil, "window-list", es.windowList_autogen, 0)
	ec.defSubr3(nil, "window-list-1", es.windowList1_autogen, 0)
	ec.defSubr1(nil, "window-live-p", es.windowLiveP_autogen, 1)
	ec.defSubr1(nil, "window-margins", es.windowMargins_autogen, 0)
	ec.defSubr1(nil, "window-minibuffer-p", es.windowMinibufferP_autogen, 0)
	ec.defSubr1(nil, "window-mode-line-height", es.windowModeLineHeight_autogen, 0)
	ec.defSubr1(nil, "window-new-normal", es.windowNewNormal_autogen, 0)
	ec.defSubr1(nil, "window-new-pixel", es.windowNewPixel_autogen, 0)
	ec.defSubr1(nil, "window-new-total", es.windowNewTotal_autogen, 0)
	ec.defSubr1(nil, "window-next-buffers", es.windowNextBuffers_autogen, 0)
	ec.defSubr1(nil, "window-next-sibling", es.windowNextSibling_autogen, 0)
	ec.defSubr2(nil, "window-normal-size", es.windowNormalSize_autogen, 0)
	ec.defSubr1(nil, "window-old-body-pixel-height", es.windowOldBodyPixelHeight_autogen, 0)
	ec.defSubr1(nil, "window-old-body-pixel-width", es.windowOldBodyPixelWidth_autogen, 0)
	ec.defSubr1(nil, "window-old-buffer", es.windowOldBuffer_autogen, 0)
	ec.defSubr1(nil, "window-old-pixel-height", es.windowOldPixelHeight_autogen, 0)
	ec.defSubr1(nil, "window-old-pixel-width", es.windowOldPixelWidth_autogen, 0)
	ec.defSubr1(nil, "window-old-point", es.windowOldPoint_autogen, 0)
	ec.defSubr2(nil, "window-parameter", es.windowParameter_autogen, 2)
	ec.defSubr1(nil, "window-parameters", es.windowParameters_autogen, 0)
	ec.defSubr1(nil, "window-parent", es.windowParent_autogen, 0)
	ec.defSubr1(nil, "window-pixel-height", es.windowPixelHeight_autogen, 0)
	ec.defSubr1(nil, "window-pixel-left", es.windowPixelLeft_autogen, 0)
	ec.defSubr1(nil, "window-pixel-top", es.windowPixelTop_autogen, 0)
	ec.defSubr1(nil, "window-pixel-width", es.windowPixelWidth_autogen, 0)
	ec.defSubr1(nil, "window-point", es.windowPoint_autogen, 0)
	ec.defSubr1(nil, "window-prev-buffers", es.windowPrevBuffers_autogen, 0)
	ec.defSubr1(nil, "window-prev-sibling", es.windowPrevSibling_autogen, 0)
	ec.defSubr2(nil, "window-resize-apply", es.windowResizeApply_autogen, 0)
	ec.defSubr2(nil, "window-resize-apply-total", es.windowResizeApplyTotal_autogen, 0)
	ec.defSubr1(nil, "window-right-divider-width", es.windowRightDividerWidth_autogen, 0)
	ec.defSubr1(nil, "window-scroll-bar-height", es.windowScrollBarHeight_autogen, 0)
	ec.defSubr1(nil, "window-scroll-bar-width", es.windowScrollBarWidth_autogen, 0)
	ec.defSubr1(nil, "window-scroll-bars", es.windowScrollBars_autogen, 0)
	ec.defSubr1(nil, "window-start", es.windowStart_autogen, 0)
	ec.defSubr1(nil, "window-tab-line-height", es.windowTabLineHeight_autogen, 0)
	ec.defSubr2(nil, "window-text-height", es.windowTextHeight_autogen, 0)
	ec.defSubr2(nil, "window-text-width", es.windowTextWidth_autogen, 0)
	ec.defSubr1(nil, "window-top-child", es.windowTopChild_autogen, 0)
	ec.defSubr1(nil, "window-top-line", es.windowTopLine_autogen, 0)
	ec.defSubr2(nil, "window-total-height", es.windowTotalHeight_autogen, 0)
	ec.defSubr2(nil, "window-total-width", es.windowTotalWidth_autogen, 0)
	ec.defSubr1(nil, "window-use-time", es.windowUseTime_autogen, 0)
	ec.defSubr1(nil, "window-valid-p", es.windowValidP_autogen, 1)
	ec.defSubr2(nil, "window-vscroll", es.windowVscroll_autogen, 0)
	ec.defSubr1(nil, "windowp", es.windowp_autogen, 1)
	ec.defSubr4(nil, "bidi-find-overridden-directionality", es.bidiFindOverriddenDirectionality_autogen, 3)
	ec.defSubr1(nil, "bidi-resolved-levels", es.bidiResolvedLevels_autogen, 0)
	ec.defSubr4(nil, "buffer-text-pixel-size", es.bufferTextPixelSize_autogen, 0)
	ec.defSubr1(nil, "current-bidi-paragraph-direction", es.currentBidiParagraphDirection_autogen, 0)
	ec.defSubr0(nil, "display--line-is-continued-p", es.displayLineIsContinuedP_autogen)
	ec.defSubr0(nil, "dump-frame-glyph-matrix", es.dumpFrameGlyphMatrix_autogen)
	ec.defSubr1(nil, "dump-glyph-matrix", es.dumpGlyphMatrix_autogen, 0)
	ec.defSubr2(nil, "dump-glyph-row", es.dumpGlyphRow_autogen, 1)
	ec.defSubr2(nil, "dump-tab-bar-row", es.dumpTabBarRow_autogen, 1)
	ec.defSubr2(nil, "dump-tool-bar-row", es.dumpToolBarRow_autogen, 1)
	ec.defSubr4(nil, "format-mode-line", es.formatModeLine_autogen, 1)
	ec.defSubr4(nil, "get-display-property", es.getDisplayProperty_autogen, 2)
	ec.defSubr1(nil, "invisible-p", es.invisibleP_autogen, 1)
	ec.defSubr0(nil, "line-pixel-height", es.linePixelHeight_autogen)
	ec.defSubr0(nil, "long-line-optimizations-p", es.longLineOptimizationsP_autogen)
	ec.defSubr3(nil, "lookup-image-map", es.lookupImageMap_autogen, 3)
	ec.defSubr1(nil, "move-point-visually", es.movePointVisually_autogen, 1)
	ec.defSubr4(nil, "set-buffer-redisplay", es.setBufferRedisplay_autogen, 4)
	ec.defSubr2(nil, "tab-bar-height", es.tabBarHeight_autogen, 0)
	ec.defSubr2(nil, "tool-bar-height", es.toolBarHeight_autogen, 0)
	ec.defSubr1(nil, "trace-redisplay", es.traceRedisplay_autogen, 0)
	ec.defSubrM(nil, "trace-to-stderr", es.traceToStderr_autogen, 1)
	ec.defSubr7(nil, "window-text-pixel-size", es.windowTextPixelSize_autogen, 0)
	ec.defSubr1(nil, "bitmap-spec-p", es.bitmapSpecP_autogen, 1)
	ec.defSubr1(nil, "clear-face-cache", es.clearFaceCache_autogen, 0)
	ec.defSubr4(nil, "color-distance", es.colorDistance_autogen, 2)
	ec.defSubr2(nil, "color-gray-p", es.colorGrayP_autogen, 1)
	ec.defSubr3(nil, "color-supported-p", es.colorSupportedP_autogen, 1)
	ec.defSubr1(nil, "color-values-from-color-spec", es.colorValuesFromColorSpec_autogen, 1)
	ec.defSubr2(nil, "display-supports-face-attributes-p", es.displaySupportsFaceAttributesP_autogen, 1)
	ec.defSubr0(nil, "dump-colors", es.dumpColors_autogen)
	ec.defSubr1(nil, "dump-face", es.dumpFace_autogen, 0)
	ec.defSubr2(nil, "face-attribute-relative-p", es.faceAttributeRelativeP_autogen, 2)
	ec.defSubr1(nil, "face-attributes-as-vector", es.faceAttributesAsVector_autogen, 1)
	ec.defSubr3(nil, "face-font", es.faceFont_autogen, 1)
	ec.defSubr1(nil, "frame--face-hash-table", es.frameFaceHashTable_autogen, 0)
	ec.defSubr4(nil, "internal-copy-lisp-face", es.internalCopyLispFace_autogen, 4)
	ec.defSubr3(nil, "internal-face-x-get-resource", es.internalFaceXGetResource_autogen, 2)
	ec.defSubr3(nil, "internal-get-lisp-face-attribute", es.internalGetLispFaceAttribute_autogen, 2)
	ec.defSubr1(nil, "internal-lisp-face-attribute-values", es.internalLispFaceAttributeValues_autogen, 1)
	ec.defSubr2(nil, "internal-lisp-face-empty-p", es.internalLispFaceEmptyP_autogen, 1)
	ec.defSubr3(nil, "internal-lisp-face-equal-p", es.internalLispFaceEqualP_autogen, 2)
	ec.defSubr2(nil, "internal-lisp-face-p", es.internalLispFaceP_autogen, 1)
	ec.defSubr2(nil, "internal-make-lisp-face", es.internalMakeLispFace_autogen, 1)
	ec.defSubr2(nil, "internal-merge-in-global-face", es.internalMergeInGlobalFace_autogen, 2)
	ec.defSubr1(nil, "internal-set-alternative-font-family-alist", es.internalSetAlternativeFontFamilyAlist_autogen, 1)
	ec.defSubr1(nil, "internal-set-alternative-font-registry-alist", es.internalSetAlternativeFontRegistryAlist_autogen, 1)
	ec.defSubr1(nil, "internal-set-font-selection-order", es.internalSetFontSelectionOrder_autogen, 1)
	ec.defSubr4(nil, "internal-set-lisp-face-attribute", es.internalSetLispFaceAttribute_autogen, 3)
	ec.defSubr4(nil, "internal-set-lisp-face-attribute-from-resource", es.internalSetLispFaceAttributeFromResource_autogen, 3)
	ec.defSubr3(nil, "merge-face-attribute", es.mergeFaceAttribute_autogen, 3)
	ec.defSubr0(nil, "show-face-resources", es.showFaceResources_autogen)
	ec.defSubr1(nil, "tty-suppress-bold-inverse-default-colors", es.ttySuppressBoldInverseDefaultColors_autogen, 1)
	ec.defSubr2(nil, "x-family-fonts", es.xFamilyFonts_autogen, 0)
	ec.defSubr5(nil, "x-list-fonts", es.xListFonts_autogen, 1)
	ec.defSubr1(nil, "x-load-color-file", es.xLoadColorFile_autogen, 1)
	ec.defSubr1(nil, "x-backspace-delete-keys-p", es.xBackspaceDeleteKeysP_autogen, 0)
	ec.defSubr6(nil, "x-begin-drag", es.xBeginDrag_autogen, 1)
	ec.defSubr1(nil, "x-display-monitor-attributes-list", es.xDisplayMonitorAttributesList_autogen, 0)
	ec.defSubr2(nil, "x-display-set-last-user-time", es.xDisplayLastUserTime_autogen, 1)
	ec.defSubr1(nil, "x-double-buffered-p", es.xDoubleBufferedP_autogen, 0)
	ec.defSubr2(nil, "x-frame-edges", es.xFrameEdges_autogen, 0)
	ec.defSubr1(nil, "x-frame-geometry", es.xFrameGeometry_autogen, 0)
	ec.defSubr1(nil, "x-frame-list-z-order", es.xFrameListZOrder_autogen, 0)
	ec.defSubr3(nil, "x-frame-restack", es.xFrameRestack_autogen, 2)
	ec.defSubr1(nil, "x-get-modifier-masks", es.xGetModifierMasks_autogen, 0)
	ec.defSubr0(nil, "x-get-page-setup", es.xGetPageSetup_autogen)
	ec.defSubr1(nil, "x-internal-focus-input-context", es.xInternalFocusInputContext_autogen, 1)
	ec.defSubr0(nil, "x-mouse-absolute-pixel-position", es.xMouseAbsolutePixelPosition_autogen)
	ec.defSubr0(nil, "x-page-setup-dialog", es.xPageSetupDialog_autogen)
	ec.defSubr1(nil, "x-print-frames-dialog", es.xPrintFramesDialog_autogen, 0)
	ec.defSubr1(nil, "x-server-input-extension-version", es.xServerInputExtensionVersion_autogen, 0)
	ec.defSubr2(nil, "x-set-mouse-absolute-pixel-position", es.xSetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr5(nil, "x-test-string-conversion", es.xTestStringConversion_autogen, 5)
	ec.defSubr6(nil, "x-translate-coordinates", es.xTranslateCoordinates_autogen, 1)
	ec.defSubr0(nil, "x-uses-old-gtk-dialog", es.xUsesOldGtkDialog_autogen)
	ec.defSubr3(nil, "x-window-property-attributes", es.xWindowPropertyAttributes_autogen, 1)
	ec.defSubr1(nil, "x-wm-set-size-hint", es.xWmSetSizeHint_autogen, 0)
	ec.defSubr0(nil, "libxml-available-p", es.libxmlAvailableP_autogen)
	ec.defSubr4(nil, "libxml-parse-html-region", es.libxmlParseHtmlRegion_autogen, 0)
	ec.defSubr4(nil, "libxml-parse-xml-region", es.libxmlParseXmlRegion_autogen, 0)
	ec.defSubr3(nil, "x-disown-selection-internal", es.xDisownSelectionInternal_autogen, 1)
	ec.defSubr2(nil, "x-get-atom-name", es.xGetAtomName_autogen, 1)
	ec.defSubr2(nil, "x-get-local-selection", es.xGetLocalSelection_autogen, 0)
	ec.defSubr4(nil, "x-get-selection-internal", es.xGetSelectionInternal_autogen, 2)
	ec.defSubr3(nil, "x-own-selection-internal", es.xOwnSelectionInternal_autogen, 2)
	ec.defSubr2(nil, "x-register-dnd-atom", es.xRegisterDndAtom_autogen, 1)
	ec.defSubr2(nil, "x-selection-exists-p", es.xSelectionExistsP_autogen, 0)
	ec.defSubr2(nil, "x-selection-owner-p", es.xSelectionOwnerP_autogen, 0)
	ec.defSubr6(nil, "x-send-client-message", es.xSendClientMessage_autogen, 6)
	ec.defSubr0(nil, "font-get-system-font", es.fontGetSystemFont_autogen)
	ec.defSubr0(nil, "font-get-system-normal-font", es.fontGetSystemNormalFont_autogen)
	ec.defSubr0(nil, "tool-bar-get-system-style", es.toolBarGetSystemStyle_autogen)
	ec.defSubr1(nil, "handle-save-session", es.handleSaveSession_autogen, 1)
	ec.defSubr1(nil, "delete-xwidget-view", es.deleteXwidgetView_autogen, 1)
	ec.defSubr1(nil, "get-buffer-xwidgets", es.getBufferXwidgets_autogen, 1)
	ec.defSubr1(nil, "kill-xwidget", es.killXwidget_autogen, 1)
	ec.defSubr7(nil, "make-xwidget", es.makeXwidget_autogen, 4)
	ec.defSubr2(nil, "set-xwidget-buffer", es.setXwidgetBuffer_autogen, 2)
	ec.defSubr2(nil, "set-xwidget-plist", es.setXwidgetPlist_autogen, 2)
	ec.defSubr2(nil, "set-xwidget-query-on-exit-flag", es.setXwidgetQueryOnExitFlag_autogen, 2)
	ec.defSubr1(nil, "xwidget-buffer", es.xwidgetBuffer_autogen, 1)
	ec.defSubr1(nil, "xwidget-info", es.xwidgetInfo_autogen, 1)
	ec.defSubr1(nil, "xwidget-live-p", es.xwidgetLiveP_autogen, 1)
	ec.defSubr3(nil, "xwidget-perform-lispy-event", es.xwidgetPerformLispyEvent_autogen, 2)
	ec.defSubr1(nil, "xwidget-plist", es.xwidgetPlist_autogen, 1)
	ec.defSubr1(nil, "xwidget-query-on-exit-flag", es.xwidgetQueryOnExitFlag_autogen, 1)
	ec.defSubr3(nil, "xwidget-resize", es.xwidgetResize_autogen, 3)
	ec.defSubr1(nil, "xwidget-size-request", es.xwidgetSizeRequest_autogen, 1)
	ec.defSubr1(nil, "xwidget-view-info", es.xwidgetViewInfo_autogen, 1)
	ec.defSubr2(nil, "xwidget-view-lookup", es.xwidgetViewLookup_autogen, 1)
	ec.defSubr1(nil, "xwidget-view-model", es.xwidgetViewModel_autogen, 1)
	ec.defSubr1(nil, "xwidget-view-p", es.xwidgetViewP_autogen, 1)
	ec.defSubr1(nil, "xwidget-view-window", es.xwidgetViewWindow_autogen, 1)
	ec.defSubr2(nil, "xwidget-webkit-back-forward-list", es.xwidgetWebkitBackForwardList_autogen, 1)
	ec.defSubr1(nil, "xwidget-webkit-estimated-load-progress", es.xwidgetWebkitEstimatedLoadProgress_autogen, 1)
	ec.defSubr3(nil, "xwidget-webkit-execute-script", es.xwidgetWebkitExecuteScript_autogen, 2)
	ec.defSubr1(nil, "xwidget-webkit-finish-search", es.xwidgetWebkitFinishSearch_autogen, 1)
	ec.defSubr2(nil, "xwidget-webkit-goto-history", es.xwidgetWebkitGotoHistory_autogen, 2)
	ec.defSubr2(nil, "xwidget-webkit-goto-uri", es.xwidgetWebkitGotoUri_autogen, 2)
	ec.defSubr3(nil, "xwidget-webkit-load-html", es.xwidgetWebkitLoadHtml_autogen, 2)
	ec.defSubr1(nil, "xwidget-webkit-next-result", es.xwidgetWebkitNextResult_autogen, 1)
	ec.defSubr1(nil, "xwidget-webkit-previous-result", es.xwidgetWebkitPreviousResult_autogen, 1)
	ec.defSubr5(nil, "xwidget-webkit-search", es.xwidgetWebkitSearch_autogen, 2)
	ec.defSubr2(nil, "xwidget-webkit-set-cookie-storage-file", es.xwidgetWebkitSetCookieStorageFile_autogen, 2)
	ec.defSubr1(nil, "xwidget-webkit-stop-loading", es.xwidgetWebkitStopLoading_autogen, 1)
	ec.defSubr1(nil, "xwidget-webkit-title", es.xwidgetWebkitTitle_autogen, 1)
	ec.defSubr1(nil, "xwidget-webkit-uri", es.xwidgetWebkitUri_autogen, 1)
	ec.defSubr2(nil, "xwidget-webkit-zoom", es.xwidgetWebkitZoom_autogen, 2)
	ec.defSubr1(nil, "xwidgetp", es.xwidgetp_autogen, 1)
}

// Subroutines count: 1663
// Constants count: 2
