'use strict';

module.exports = function(grunt) {
  var args = [
    ['--base-font-size', '10'],
    ['--authors', 'Seth Vincent'],
    ['--publisher', 'seattle.io'],
    ['--extra-css', 'test.css']
  ];

  // Project configuration.
  grunt.initConfig({

    markdownpdf: {
      options: {},
      files: {
        src: "*.md",
        dest: "dest"
      }
    }

  });

  grunt.loadNpmTasks('grunt-markdown-pdf');

  // Default task.
  grunt.registerTask('default', ['markdownpdf']);

};
